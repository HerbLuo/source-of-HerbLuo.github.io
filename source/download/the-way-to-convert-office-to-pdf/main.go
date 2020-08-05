package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

var wordDir, _ = filepath.Abs("./")
var vbsFilepath = filepath.Join(wordDir, "excel2pdf.vbs")
var randGen = rand.New(rand.NewSource(time.Now().UnixNano()))

func convert2pdf(w http.ResponseWriter, r *http.Request) {
	// 生成临时文件名与路径
	filename := strconv.Itoa(randGen.Int())
	fmt.Println("handle a request, the temp filename is ", filename)
	extension := ".xlsx"
	reqPath := r.URL.Path
	if reqPath == "/docx-to-pdf/" {
		extension = ".docx"
	}
	excelFilepath := filepath.Join(wordDir, filename + extension)
	pdfFilepath := filepath.Join(wordDir, filename + ".pdf")
	excelFile, err := os.Create(excelFilepath)
	if err != nil {
		panic(err)
	}
	// 保存请求体到文件
	defer r.Body.Close()
	if _, err = io.Copy(excelFile, r.Body); err != nil {
		panic(err)
	}
	if err = excelFile.Close(); err != nil { // 必须同步关闭该文件，不然 vbs 程序运行时无法打开该文件
		panic(err)
	}
	// 转换
	c := exec.Command("cmd", "/K", vbsFilepath, excelFilepath)
	if err = c.Run(); err != nil {
		panic(err)
	}
	// 返回生成的pdf文件
	pdfFile, err := os.Open(pdfFilepath)
	defer pdfFile.Close()
	if _, err = io.Copy(w, pdfFile); err != nil {
		panic(err)
	}
	go func() {
		if err := os.Remove(excelFilepath); err != nil {
			fmt.Println(err)
		}
		if err := os.Remove(pdfFilepath); err != nil {
			fmt.Println(err)
		}
	}()
}

func createVbsScriptFile() {
	f, err := os.Create(vbsFilepath)
	if err != nil {
		panic(err)
	}
	if _, err = f.Write([]byte(xls2pdfScript)); err != nil {
		panic(err)
	}
	if err = f.Close(); err != nil {
		panic(err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte(HomePageHtml))
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/xlsx-to-pdf/", convert2pdf)
	http.HandleFunc("/docx-to-pdf/", convert2pdf)
	fmt.Println("listen and served on port 9090")
	createVbsScriptFile()
	panic(http.ListenAndServe(":9090", nil))
}

const xls2pdfScript = `Option Explicit

Sub main()
Dim Filepath, FilepathList, Extension
Filepath = WScript.Arguments(0)
FilepathList = split(Filepath, ".")
Extension = FilepathList(ubound(FilepathList))
if Extension = "docx" Then
  convertDocToPDF Filepath
end if
if Extension = "xlsx" Then
  convertXlsToPDF Filepath
end if
End Sub

Function convertDocToPDF(Path)
Dim Objshell, ParentFolder, BaseName, docApp, Doc, PDFPath
Set Objshell = CreateObject("scripting.filesystemobject")
ParentFolder = Objshell.GetParentFolderName(Path)
BaseName = Objshell.GetBaseName(Path)
PDFPath = parentFolder & "\" & BaseName & ".pdf"

Set docApp = CreateObject("Word.application")
docApp.Visible = False
Set Doc = docApp.Documents.Open(Path)

Doc.SaveAs PDFPath, 17
Doc.saved = True
Doc.close
docApp.quit

Set Objshell = Nothing
End Function

Function convertXlsToPDF(Path)
Dim Objshell, ParentFolder, BaseName, XlsApp, Doc, PDFPath
Set Objshell = CreateObject("scripting.filesystemobject")
ParentFolder = Objshell.GetParentFolderName(Path)
BaseName = Objshell.GetBaseName(Path)
PDFPath = parentFolder & "\" & BaseName & ".pdf"

Set XlsApp = CreateObject("Excel.application")
Set Doc = XlsApp.Workbooks.Open(Path)
XlsApp.CalculateFull
Doc.ExportAsFixedFormat 0, PDFPath
Doc.saved = True
Doc.close
XlsApp.quit

Set Objshell = Nothing
End Function

Call main
`

const HomePageHtml  = `
<!DOCTYPE html>
<body><input type="file" onchange="inputChange(event)" /></body>
<script>
    let i = 0;
    async function inputChange(e) {
        const file = e.target.files[0];
        const url =  file.name.split(".").slice(-1)[0].startsWith("doc")
            ? "http://localhost:9090/docx-to-pdf/"
            : "http://localhost:9090/xlsx-to-pdf/";
        const r = await fetch(url, { method: "POST", body: file });
        const a = document.createElement('a');
        a.innerHTML = (++i + ".pdf").padStart(7, "0");
        a.target = "_blank"; a.style.display = "block";
        a.href = URL.createObjectURL(await r.blob());
        document.body.appendChild(a);
    }
</script></html>
`

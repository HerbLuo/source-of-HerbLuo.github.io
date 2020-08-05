---
title: xlsx, docx转pdf
date: 2019-05-25 00:00:00
tags: [the-way, office, pdf]
---

`Office` 如何转`PDF`？
网上搜罗了一下，方案很多。但要做到和`Microsoft Office`打印出来的一模一样。还是得靠`Windows`和`Office`才行。

这里提供一个`EXE`文件(2.7M)，双击运行即可开放9090端口（http协议），直接请求即可完成转换。

<!-- more -->

## 使用方法

1. 在一台安装了`Office 2007`以上版本的`Windows`电脑。
双击运行下载的`EXE`文件(下载地址位于附录中)，这样，一个`office`转`PDF`服务器就做好了
(弹出的防火墙弹窗需通过，不然只能本地访问)
![](/images/the-way-to-convert-xlsx-to-pdf-0.png)

2. 访问[http://localhost:9090/](http://localhost:9090/)
选择一个`xlsx`或`docx`文件，稍等片刻，即可完成转换。
![](/images/the-way-to-convert-xlsx-to-pdf-1.png)

3. 接口描述
  3.1. xlsx转pdf
```
url: http://{host}:9090/xlsx-to-pdf/
method: POST
body: 整个office源文件（非form形式）
```

  3.2. docx转pdf
```
url: http://{host}:9090/docx-to-pdf/
method: POST
body: 整个office源文件（非form形式）
```

## 实现原理

将下面一段脚本另存为`foo.vbs`

```vbs
Option Explicit

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
```

将任意`docx`或`xlsx`文件拖至上述文件`foo.vbs`上方。稍等片刻，脚本即会生成相应的`PDF`文件（位于源office文件相同目录）。

好了，剩下的事情就是开放一个`http`服务并调用上述脚本即可。
这里用的是`go lang`（因为相对熟悉些，更重要的是部署方便）
源码可参考附录，可自行编译。

## 更好的方案

就我个人而言，实现一个功能，更喜欢透明度更高的脚本而不是用二进制文件这种“乱七八糟”的东西。所以，更希望使用`powershell`脚本创建`http`服务并调用`Office`的接口。不过目前来看时间有限，等下次遇到相关需求时再尝试吧。

## 附录：
office-to-pdf.exe
[64位下载](/download/the-way-to-convert-office-to-pdf/office-to-pdf-amd64.exe) 
```
SHA256=106ba9b780f626b5939a7ac364c52d053809201d19549098e5ff2c8fff4fb581
MD5=dcee8bfa6783a4ce2470a8b500aeb01d
```
[32位下载](/download/the-way-to-convert-office-to-pdf/office-to-pdf-i386.exe)
```
SHA256=f2db1e21058a1b40348b0e9b951c09931217b779947b6c0bf2daaa30aa84a06c
MD5=c76789b72cc8c5e77df63ac71459cbab
```
[GO源码](/download/the-way-to-convert-office-to-pdf/main.go)

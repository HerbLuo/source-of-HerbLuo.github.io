package cn.cloudself.components.filter;

import javax.servlet.*;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.util.regex.Pattern;

/**
 * @author HerbLuo
 * @version 1.0.0.d
 */
public class CORSFilter implements Filter {

    // 使用时，替换此处即可
    private final static String CLOUDSELF_CN = "http://shop.cloudself.cn";

    private final static String ACCESS_CONTROL_ALLOW_ORIGIN = "Access-Control-Allow-Origin";
    private final static Pattern patternLocalhost = Pattern.compile("https?://localhost(:[0-9]{1,5})?");
    private final static Pattern pattern127 = Pattern.compile("https?://127.[0-9]{0,3}.[0-9]{0,3}.[0-9]{0,3}(:[0-9]{1,5})?");
    private final static Pattern pattern192_168 = Pattern.compile("https?://192.168.[0-9]{0,3}.[0-9]{0,3}(:[0-9]{1,5})?");
    private final static Pattern pattern172_16 = Pattern.compile("https?://172.(1[6-9]|2[0-9]|3[01]).[0-9]{0,3}.[0-9]{0,3}(:[0-9]{1,5})?");
    private final static Pattern pattern10 = Pattern.compile("https?://10.[0-9]{0,3}.[0-9]{0,3}.[0-9]{0,3}(:[0-9]{1,5})?");

    @Override
    public void init(FilterConfig filterConfig) throws ServletException {
    }

    @Override
    public void doFilter(ServletRequest request, ServletResponse response,
                         FilterChain chain) throws IOException, ServletException {
        assert request instanceof HttpServletRequest : "该CORSFilter  只支持 http协议";
        assert response instanceof HttpServletResponse : "该CORSFilter 只支持 http协议";

        HttpServletRequest req = (HttpServletRequest) request;
        HttpServletResponse res = (HttpServletResponse) response;

        String origin = req.getHeader("Origin");

        if (CLOUDSELF_CN.equals(origin)) {
            res.setHeader(ACCESS_CONTROL_ALLOW_ORIGIN, CLOUDSELF_CN);
        } else if (
                patternLocalhost.matcher(origin).matches()
                        || pattern127.matcher(origin).matches()
                        || pattern192_168.matcher(origin).matches()
                        || pattern172_16.matcher(origin).matches()
                        || pattern10.matcher(origin).matches()
                ) {
            res.setHeader(ACCESS_CONTROL_ALLOW_ORIGIN, origin);
        } else {
            return;
        }

        if ("OPTIONS".equals(req.getMethod())) {
            res.setHeader("Access-Control-Allow-Headers", "content-type");
            res.setHeader("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS, DELETE"); //允许的请求方法
            res.setHeader("Access-Control-Max-Age", "2592000"); //options 请求允许缓存30天
        }
        chain.doFilter(request, response);

    }

    @Override
    public void destroy() {
    }
}

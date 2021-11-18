package ieetubase.common.filter;

import org.apache.http.HttpStatus;
import org.springframework.web.filter.OncePerRequestFilter;

import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

public class HttpPreProcessFilter extends OncePerRequestFilter {
    @Override
    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain filterChain) throws ServletException, IOException {
        String method = request.getMethod();
        if(method.equalsIgnoreCase("OPTIONS")) {
            response.setStatus(HttpStatus.SC_OK);
            return;
        }
        if(!(method.equalsIgnoreCase("GET") || method.equalsIgnoreCase("POST"))) {
            response.sendError(HttpServletResponse.SC_METHOD_NOT_ALLOWED);

            return;
        }
        filterChain.doFilter(request, response);
    }
}

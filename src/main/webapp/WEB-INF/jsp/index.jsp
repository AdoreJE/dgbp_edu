<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<%@ taglib prefix="spring" uri="http://www.springframework.org/tags" %>
<!DOCTYPE html>
<html>
    <head>
        <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
        <script src="${pageContext.request.contextPath}/js/index.js"></script>
        <link rel="stylesheet" type="text/css" href="${pageContext.request.contextPath}/css/style.css">
        <style>
          
        </style>
    </head>
    <body>
        <h1>대구시 블록체인 활용 공공서비스</h1>
        <img src="${pageContext.request.contextPath}/images/main_visual_02.png">
        <br>
        <input class="textBox top" type="text" placeholder="값을 입력해주세요" id="value">
        <input class="buttonBox top btn-blue" type="button" onclick="invoke()" value="invoke">
        <br>
        <input class="textBox bottom" type="text" placeholder="번호를 입력해주세요" id="number">
        <input class="buttonBox bottom" type="button" onclick="query()" value="query">
        <div id="loading">
            <img src="${pageContext.request.contextPath}/images/loading.gif" width="300" height="100">
        </div>
    </body>
</html>
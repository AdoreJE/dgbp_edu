$(document).ready(function(){
    $('#loading').hide();
})
.ajaxStart(function(){
    $('#loading').show();
})
.ajaxStop(function(){
    $('#loading').hide();
})
function invoke() {
    let value = document.getElementById("value").value;
    if (value == "" || value == undefined){
        alert("입력해주세요");
        return;
    }
    $.ajax({
        type:"POST",
           url : "https://dgbp.daeguedu.com/dgbp/api/edu/invoke.do",
           data: getJSONData("ping", value),
           dataType : 'json',
           contentType : "application/json; charset=UTF-8",
           
           success: function(data) {
               console.log(data)
               if (data.status === "invalid") {
                    alert(data.message)
               } else {
                    alert("당신은 " + data.response.payload + "번째로 invoke 함수를 사용하셨습니다.")
               }
           },
           error: function(error){
               
           }
       });  
}
function query() {
    let num = document.getElementById("number").value;
    if (num === "" || num == undefined){
        alert("입력해주세요");
        return;
    }
    $.ajax({
        type:"POST",
           url : "https://dgbp.daeguedu.com/dgbp/api/edu/query.do",
           data: getJSONData("pong", num),
           dataType : 'json',
           contentType : "application/json; charset=UTF-8",

           success: function(data) {
               console.log(data)
               if (data.status === "invalid") {
                    alert(data.message)
               } else if(data.response.payload === undefined) {
                   alert("올바른 번호를 입력해주세요")
               } else {
                    alert(num + "번째 값은 " + data.response.payload + " 입니다")
               }
           },
           error: function(error){
               
           }
       });  
}

// 체인코드 함수 호출 부분
function getJSONData(fnName, arg1){
	let data = {
		"chaincodeName" : "eduBook", // 체인코드 이름
		"functionName" : fnName, // 함수 이름
		"args" : [arg1] // 함수 매개변수
	}
	return JSON.stringify(data);
}
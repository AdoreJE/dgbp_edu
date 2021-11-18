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

function getJSONData(fnName, args){
    return '{"apiKey":"","chaincodeName":"eduBook","functionName":"'+ fnName + '","args":["' + args + '"]}';
}
function addCustomer() {
  let t2f = false;
  $("input").each(function() {
    t2f = checkInput(this);
    if (t2f) {
      return false;
    }
  });
  if (t2f) {
    alert('所有内容必须都填，请检查后重新提交');
		return;
  }
  var postStr;
  postStr= 'operation=' + encodeURIComponent('insert');
  postStr= postStr + '&cstmname='+encodeURIComponent(document.getElementById("cstmname").value.replace(/\s+/g,""));
  postStr= postStr + '&cstmtype='+encodeURIComponent(document.getElementById("cstmtype").value.replace(/\s+/g,""));
  postStr= postStr + '&city='+encodeURIComponent(document.getElementById("city").value.replace(/\s+/g,""));
  postStr= postStr + '&area='+encodeURIComponent(document.getElementById("area").value.replace(/\s+/g,""));
  postStr= postStr + '&address='+encodeURIComponent(document.getElementById("address").value.replace(/\s+/g,""));
  postStr= postStr + '&owner_name='+encodeURIComponent(document.getElementById("owner_name").value.replace(/\s+/g,""));
  postStr= postStr + '&telephone='+encodeURIComponent(document.getElementById("telephone").value.replace(/\s+/g,""));
  postStr= postStr + '&police='+encodeURIComponent(document.getElementById("police").value.replace(/\s+/g,""));
  postStr= postStr + '&axis='+encodeURIComponent(document.getElementById("axis").value.replace(/\s+/g,""));
  postStr= postStr + '&remark='+encodeURIComponent(document.getElementById("remark").value.replace(/\s+/g,""));


  var xhr = new XMLHttpRequest();
  
  xhr.onload = function () {
    if (xhr.status == 200) {
      $('#rsts').html(xhr.responseText);
    }
  };
  console.log(postStr);
  xhr.open("POST", "/handle_customers", true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  xhr.send(postStr);
};



var el = document.getElementById('customer_add');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    addCustomer(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    addCustomer(e);
  });
}
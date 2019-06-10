function addPayment() {
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
  postStr= postStr + '&create_date='+encodeURIComponent(document.getElementById("create_date").value.replace(/\s+/g,""));
  postStr= postStr + '&amount='+encodeURIComponent(document.getElementById("amount").value.replace(/\s+/g,""));
  postStr= postStr + '&pmtype='+encodeURIComponent(document.getElementById("pmtype").value.replace(/\s+/g,""));
  postStr= postStr + '&remark='+encodeURIComponent(document.getElementById("remark").value.replace(/\s+/g,""));


  var xhr = new XMLHttpRequest();
  
  xhr.onload = function () {
    if (xhr.status == 200) {
      $('#rsts').html(xhr.responseText);
    }
  };
  
  xhr.open("POST", "/handle_payments", true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  xhr.send(postStr);
  console.log(postStr);
};

let cstmSelects = getCustomerName();
$("#cstmname").autocomplete({
  source: cstmSelects
});

$("#create_date").datepicker();

var el = document.getElementById('payment_add');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    addPayment(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    addPayment(e);
  });
}

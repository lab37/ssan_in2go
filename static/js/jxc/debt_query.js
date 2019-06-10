function debt_query(e) {

  let tgt = getTarget(e);
  var postStr;
  postStr = 'operation=' + encodeURIComponent('select');
  if (tgt.id == 'debt_query_all') {

    postStr = postStr + '&tt=' + encodeURIComponent('all');
  } else {
    let t2f = true;
    $(".input").each(function () {
      t2f = checkInput(this);
      if (!t2f) {
        return false;
      }
    });
    if (t2f) {
      alert('请至少填写一项，查询全部请直接点击查询全部按钮');
      return;
    }
    postStr = postStr + '&tt=' + encodeURIComponent('some');
    postStr = postStr + '&cstmname=' + encodeURIComponent(document.getElementById("cstmname").value.replace(/\s+/g, ""));
  }



  var xhr = new XMLHttpRequest();

  xhr.onload = function () {
    if (xhr.status == 200) {

      responseObject = JSON.parse(xhr.responseText);
      var newContent = '';
      for (var i = 0; i < responseObject.length; i++) {    // Loop through object
        newContent += '<tr id="' + responseObject[i].Id + '">';
        newContent += '<td name="id">' + responseObject[i].Id + '</td>';
        newContent += '<td name="cstmname">' + responseObject[i].CstmName + '</td>';
        newContent += '<td name="amount">' + responseObject[i].Amount + '</td>';
        newContent += '<td name="remark">' + responseObject[i].Remark + '</td>';
        newContent += '</tr>';
      }
      // Update the page with the new content
      document.getElementById('rows').innerHTML = newContent;
      $('#debt_results').show();
      $('#export').show()
    }
  };

  xhr.open("POST", "/handle_debts", true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  xhr.send(postStr);
  console.log(postStr);
};
$('#debt_results').hide();
$('#export').hide();
$('#export').click(function(){
  table2xlsx('xlsx','debt_results');
});
$("#pmdate").datepicker();

var el = document.getElementById('debt_query_some');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    debt_query(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    debt_query(e);
  });
}

var el2 = document.getElementById('debt_query_all');
if (el2.addEventListener) {
  el2.addEventListener('click', function (e) {
    debt_query(e);
  }, false);
} else {
  el2.attachEvent('onclick', function (e) {
    debt_query(e);
  });
}


let cstmSelects = getCustomerName();
$("#cstmname").autocomplete({
  source: cstmSelects
});
$("#srcname").autocomplete({
  source: cstmSelects
});

$("#change_table").hide();
function hd() {
  $("#change_table").hide();
}
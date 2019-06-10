function stockQuery(e) {

  let tgt = getTarget(e);
  var postStr;
  postStr = 'operation=' + encodeURIComponent('select');
  if (tgt.id == 'stock_query_all') {
    postStr = postStr + '&tt=' + encodeURIComponent('all');
  } else {
    let t2f = true;
    $("input").each(function () {
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
    postStr = postStr + '&prdtname=' + encodeURIComponent(document.getElementById("prdtname").value.replace(/\s+/g, ""));
    postStr = postStr + '&specific=' + encodeURIComponent(document.getElementById("specific").value.replace(/\s+/g, ""));
    postStr = postStr + '&mac=' + encodeURIComponent(document.getElementById("mac").value.replace(/\s+/g, ""));
  }

  var xhr = new XMLHttpRequest();

  xhr.onload = function () {
    if (xhr.status == 200) {

      responseObject = JSON.parse(xhr.responseText);
      var newContent = '';
      
      for (var i = 0; i < responseObject.length; i++) {    // Loop through object
        newContent += '<tr id="' + responseObject[i].Id + '">';
        newContent += '<td name="id">' + responseObject[i].Id + '</td>';
        newContent += '<td name="prdtname">' + responseObject[i].PrdtName + '</td>';
        newContent += '<td name="specific">' + responseObject[i].Specific + '</td>';
        newContent += '<td name="mac">' + responseObject[i].Mac + '</td>';
        newContent += '<td name="sn">' + responseObject[i].Sn + '</td>';
        newContent += '<td name="quantity">' + responseObject[i].Quantity + '</td>';

        newContent += '<td name="remark">' + responseObject[i].Remark + '</td>';
        newContent += '</tr>';
      }
      // Update the page with the new content
      document.getElementById('rows').innerHTML = newContent;
      $('#stocks_results').show();
      $('#export').show();
    }
  };

  xhr.open("POST", "/handle_stocks", true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  xhr.send(postStr);
  console.log(postStr);
};
$('#stocks_results').hide();
$('#export').hide();
$('#export').click(function(){
  table2xlsx('xlsx','stocks_results');
});

var el = document.getElementById('stock_query_some');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    stockQuery(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    stockQuery(e);
  });
}

var el2 = document.getElementById('stock_query_all');
if (el2.addEventListener) {
  el2.addEventListener('click', function (e) {
    stockQuery(e);
  }, false);
} else {
  el2.attachEvent('onclick', function (e) {
    stockQuery(e);
  });
}
function onWayProductQuery(e) {

  let tgt = getTarget(e);
  var postStr;
  postStr = 'operation=' + encodeURIComponent('select');
  postStr += '&vector=0';
  if (tgt.id == 'onway_product_query_all') {
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
    postStr = postStr + '&prdtname=' + encodeURIComponent(document.getElementById("prdtname").value.replace(/\s+/g, ""));
    postStr = postStr + '&specific=' + encodeURIComponent(document.getElementById("specific").value.replace(/\s+/g, ""));
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
        newContent += '<td name="prdtname">' + responseObject[i].PrdtName + '</td>';
        newContent += '<td name="specific">' + responseObject[i].Specific + '</td>';
        newContent += '<td name="quantity">' + responseObject[i].Quantity + '</td>';
        newContent += '<td name="remark">' +responseObject[i].Remark + '</td>';

        newContent += '</tr>';
      }
      // Update the page with the new content
      document.getElementById('rows').innerHTML = newContent;
      $('#onway_product_results').show();
      $('#export').show();
    }
  };

  xhr.open("POST", "/handle_onway_products", true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  xhr.send(postStr);
  console.log(postStr);
};
$('#onway_product_results').hide();
$('#export').hide();
$('#export').click(function(){
  table2xlsx('xlsx','onway_product_results');
});
let cstmSelects = getCustomerName();
$("#cstmname").autocomplete({
  source: cstmSelects
});


var el = document.getElementById('onway_product_query_some');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    onWayProductQuery(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    onWayProductQuery(e);
  });
}

var el2 = document.getElementById('onway_product_query_all');
if (el2.addEventListener) {
  el2.addEventListener('click', function (e) {
    onWayProductQuery(e);
  }, false);
} else {
  el2.attachEvent('onclick', function (e) {
    onWayProductQuery(e);
  });
}
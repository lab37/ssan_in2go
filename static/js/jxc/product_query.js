function product_query(e) {

  let tgt = getTarget(e);
  var postStr;
  if (tgt.id == 'product_query_all') {
    postStr = 'operation=' + encodeURIComponent('select');
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
    postStr = 'operation=' + encodeURIComponent('select');
    postStr = postStr + '&tt=' + encodeURIComponent('some');
    postStr = postStr + '&specific=' + encodeURIComponent(document.getElementById("specific").value.replace(/\s+/g, ""));
    postStr = postStr + '&prdtname=' + encodeURIComponent(document.getElementById("prdtname").value.replace(/\s+/g, ""));
    postStr = postStr + '&inventor=' + encodeURIComponent(document.getElementById("inventor").value.replace(/\s+/g, ""));
  }
  console.log(postStr);
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
        newContent += '<td name="inventor">' + responseObject[i].Inventor + '</td>';
        newContent += '<td name="unit">' + responseObject[i].Unit + '</td>';
        newContent += '<td name="ivtype">' + responseObject[i].IvType + '</td>';
        newContent += '<td name="remark">' + responseObject[i].Remark + '</td>';
        newContent += '<td>' + '<input type="button" class="chg" onclick="upt(' + "'" + responseObject[i].Id + "'" + ')" value="修改" /> <input type="button" class="del" onclick="del(' + "'" + responseObject[i].Id + "'" + ')" value="删除" />' + '</td>';
        newContent += '</tr>';
      }
      // Update the page with the new content
      document.getElementById('rows').innerHTML = newContent;
      $('#product_results').show();
      $('#export').show();
    }
  };

  xhr.open("POST", "/handle_products", true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  xhr.send(postStr);
};
$('#product_results').hide();
$('#export').hide();
$('#export').click(function () {
  table2xlsx('xlsx', 'product_results');
});


let prdtObjects = getProductNS();
let prdtNames = new Array();
for (var i = 0; i < prdtObjects.length; i++) {
  prdtNames.push(prdtObjects[i].PrdtName);
}
$("#prdtname").autocomplete({
  source: prdtNames
});

$("#specific").focus(function () {
  let prdtSpecifics = new Array();
  
  for (var i = 0; i < prdtObjects.length; i++) {
    if (prdtObjects[i].PrdtName == $("#prdtname").val()) {
      prdtSpecifics.push(prdtObjects[i].Specific);
    }
  }
  $("#specific").autocomplete({
    source: prdtSpecifics
  });
});


var el = document.getElementById('product_query_some');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    product_query(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    product_query(e);
  });
}

var el2 = document.getElementById('product_query_all');
if (el2.addEventListener) {
  el2.addEventListener('click', function (e) {
    product_query(e);
  }, false);
} else {
  el2.attachEvent('onclick', function (e) {
    product_query(e);
  });
}

function upt(id) {
  let domStr = '<legend><span>修改记录</span></legend>';
  $('#' + id).children().each(function (index, element) {
    if ($(element).attr("name") != undefined) {
      domStr += '<label>' + index + '：</label><input id="update_' + $(element).attr("name") + '" value="' + $(element).text() + '">'
    }
  });
  domStr += '<input type="button" class="sv" onclick="sv(' + "'" + id + "'" + ')" value="保存" /><input type="button" class="sv" onclick="hd()" value="取消" />';
  $("#change_table").html(domStr);
  $("#change_table").show();
  // window.location.hash = "#change_table";
  $("html,body").animate({ scrollTop: $("#change_table").offset().top }, 1000);

}

function sv(id) {
  $.post("handle_products", {
    operation: 'update',
    id: id,
    prdtname: document.getElementById("update_prdtname").value.replace(/\s+/g, ""),
    specific: document.getElementById("update_specific").value.replace(/\s+/g, ""),
    inventor: document.getElementById("update_inventor").value.replace(/\s+/g, ""),
    unit: document.getElementById("update_unit").value.replace(/\s+/g, ""),
    ivtype: document.getElementById("update_ivtype").value.replace(/\s+/g, ""),
    remark: document.getElementById("update_remark").value.replace(/\s+/g, "")
  },
    function (data, status) {
      alert(data);
      if (status == "success") {
        $("#change_table").hide();
      }
    });
}

$("#change_table").hide();
function hd() {
  $("#change_table").hide();
}

function del(id) {
  var cfm = confirm("确认要删这条记录吗？");
  if (cfm == true) {
    $.post("handle_products",
      {
        operation: "delete",
        id: id
      },
      function (data, status) {
        if (status == "success") {
          alert(data);
          $("#" + id).hide();

        } else {
          alert("服务器错误，请联系周京成");
        }
      });
  }
  else {
    return;
  }
}
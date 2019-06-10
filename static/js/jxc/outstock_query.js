function income_query(e) {
  
  let tgt = getTarget(e);
  var postStr;
  if (tgt.id == 'outstock_query_all') {
    postStr = 'operation=' + encodeURIComponent('select');
    postStr = postStr + '&tt=' + encodeURIComponent('all');
  } else {
    let t2f = true;
  $(".input").each(function() {
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
    postStr = postStr + '&cstmname=' + encodeURIComponent(document.getElementById("cstmname").value.replace(/\s+/g,""));
    postStr = postStr + '&create_date=' + encodeURIComponent(document.getElementById("create_date").value.replace(/\s+/g,""));
    postStr = postStr + '&prdtname=' + encodeURIComponent(document.getElementById("prdtname").value.replace(/\s+/g,""));
    postStr = postStr + '&specific=' + encodeURIComponent(document.getElementById("specific").value.replace(/\s+/g,""));
    postStr = postStr + '&mac=' + encodeURIComponent(document.getElementById("mac").value.replace(/\s+/g,""));
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
        newContent += '<td name="mac">' + responseObject[i].Mac + '</td>';
        newContent += '<td name="sn">' + responseObject[i].Sn + '</td>';
        newContent += '<td name="create_date">' + responseObject[i].CreateDate + '</td>';
        newContent += '<td name="quantity">' + responseObject[i].Quantity + '</td>';
        newContent += '<td name="remark">' + responseObject[i].Remark + '</td>';
        newContent += '<td>' + '<input type="button" class="chg" onclick="upt(' +"'"+ responseObject[i].Id +"'"+ ')" value="修改" /> <input type="button" class="del" onclick="del(' +"'"+ responseObject[i].Id +"'"+ ')" value="删除" />' + '</td>';
        newContent += '</tr>';
      }
      // Update the page with the new content
      document.getElementById('rows').innerHTML = newContent;
      $('#outstock_results').show();
      $('#export').show()
    }
  };

  xhr.open("POST", "/handle_outstocks", true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  xhr.send(postStr);
  console.log(postStr);
};
$('#outstock_results').hide();
$('#export').hide();
$('#export').click(function(){
  table2xlsx('xlsx','income_results');
});
$("#create_date").datepicker();

let cstmSelects = getCustomerName();
$("#cstmname").autocomplete({
  source: cstmSelects
});


var el = document.getElementById('outstock_query_some');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    income_query(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    income_query(e);
  });
}

var el2 = document.getElementById('outstock_query_all');
if (el2.addEventListener) {
  el2.addEventListener('click', function (e) {
    income_query(e);
  }, false);
} else {
  el2.attachEvent('onclick', function (e) {
    income_query(e);
  });
}

function upt(id) {
  let domStr='<legend><span>修改记录</span></legend>';
  $('#' + id).children().each(function (index, element) {
    if($(element).attr("name") != undefined){
    domStr += '<label>'+index +'：</label><input id="update_' + $(element).attr("name") + '" value="' + $(element).text() + '">'
    }
  });
  domStr += '<input type="button" class="sv" onclick="sv(' +"'"+ id +"'"+ ')" value="保存" /><input type="button" class="sv" onclick="hd()" value="取消" />';
  console.log(domStr);
  $("#change_table").html(domStr);
  $("#change_table").show();
  $("html,body").animate({scrollTop:$("#change_table").offset().top},1000);
}

function sv(id) {
  $.post("handle_outstocks",{
  operation:'update',
  id:id,
  cstmname : document.getElementById("update_cstmname").value.replace(/\s+/g,""),
  prdtname : document.getElementById("update_prdtname").value.replace(/\s+/g,""),
  specific : document.getElementById("update_specific").value.replace(/\s+/g,""),
  mac : document.getElementById("update_mac").value.replace(/\s+/g,""),
  sn : document.getElementById("update_sn").value.replace(/\s+/g,""),
  create_date : document.getElementById("update_create_date").value.replace(/\s+/g,""),
  quantity : document.getElementById("update_quantity").value.replace(/\s+/g,""),
  remark : document.getElementById("update_remark").value.replace(/\s+/g,"")
  },
  function(data,status){
    alert(data);
    if(status=="success"){
      $("#change_table").hide();
    }
  });
}

$("#change_table").hide();
function hd(){
  $("#change_table").hide();
  }

function del(id) {
  var cfm = confirm("确认要删这条记录吗？");
  if (cfm == true) {
    $.post("handle_outstocks",
      {
        operation: "delete",
        id: id
      },
      function (data, status) {
        if (status == "success") {
          alert(data);
          $("#" +id).hide();
        } else {
          alert("服务器错误，请联系周京成");
        }
      });
  }
  else {
    return;
  }
}
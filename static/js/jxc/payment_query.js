function payment_query(e) {
  
  let tgt = getTarget(e);
  var postStr;
  if (tgt.id == 'payment_query_all') {
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
    postStr = postStr + '&create_date=' + encodeURIComponent(document.getElementById("create_date").value.replace(/\s+/g,""));
    postStr = postStr + '&cstmname=' + encodeURIComponent(document.getElementById("cstmname").value.replace(/\s+/g,""));
  }

  

  var xhr = new XMLHttpRequest();

  xhr.onload = function () {
    if (xhr.status == 200) {

      responseObject = JSON.parse(xhr.responseText);
      var newContent = '';
      
      for (var i = 0; i < responseObject.length; i++) {    // Loop through object
        newContent += '<tr id="' + responseObject[i].Id + '">';
        newContent += '<td name="cstmname">' + responseObject[i].CstmName + '</td>';
        newContent += '<td name="create_date">' + responseObject[i].CreateDate + '</td>';
        newContent += '<td name="pmtype">' + responseObject[i].PmType + '</td>';
        newContent += '<td name="amount">' + responseObject[i].Amount + '</td>';
        newContent += '<td name="remark">' + responseObject[i].Remark + '</td>';
        newContent += '<td>' + '<input type="button" class="chg" onclick="upt(' +"'"+ responseObject[i].Id +"'"+ ')" value="修改" /> <input type="button" class="del" onclick="del(' +"'"+ responseObject[i].Id +"'"+ ')" value="删除" />' + '</td>';
        newContent += '</tr>';
      }
      // Update the page with the new content
      document.getElementById('rows').innerHTML = newContent;
      $('#payment_results').show();
      $('#export').show();
    }
  };

  xhr.open("POST", "/handle_payments", true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  xhr.send(postStr);
  console.log(postStr);
};

$('#payment_results').hide();
$('#export').hide();
$('#export').click(function(){
  table2xlsx('xlsx','payment_results');
});

$("#create_date").datepicker();

let cstmSelects = getCustomerName();
$("#cstmname").autocomplete({
  source: cstmSelects
});

var el = document.getElementById('payment_query_some');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    payment_query(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    payment_query(e);
  });
}

var el2 = document.getElementById('payment_query_all');
if (el2.addEventListener) {
  el2.addEventListener('click', function (e) {
    payment_query(e);
  }, false);
} else {
  el2.attachEvent('onclick', function (e) {
    payment_query(e);
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
  $("#change_table").html(domStr);
  $("#change_table").show();
  $("html,body").animate({scrollTop:$("#change_table").offset().top},1000);
}

function sv(id) {
  $.post("handle_payments",{
  operation:'update',
  id:id,
  cstmname : document.getElementById("update_cstmname").value.replace(/\s+/g,""),
  create_date : document.getElementById("update_create_date").value.replace(/\s+/g,""),
  pmtype : document.getElementById("update_pmtype").value.replace(/\s+/g,""),
  amount : document.getElementById("update_amount").value.replace(/\s+/g,""),
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
    $.post("handle_payments",
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
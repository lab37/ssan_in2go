function customer_query(e) {
  
  let tgt = getTarget(e);
  var postStr;
  if (tgt.id == 'customer_query_all') {
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
  postStr= postStr + '&cstmname='+encodeURIComponent(document.getElementById("cstmname").value.replace(/\s+/g,""));
  postStr= postStr + '&city='+encodeURIComponent(document.getElementById("city").value.replace(/\s+/g,""));
  postStr= postStr + '&area='+encodeURIComponent(document.getElementById("area").value.replace(/\s+/g,""));
  postStr= postStr + '&police='+encodeURIComponent(document.getElementById("police").value.replace(/\s+/g,""));
  }

  var xhr = new XMLHttpRequest();

  xhr.onload = function () {
    if (xhr.status == 200) {
      responseObject = JSON.parse(xhr.responseText);
      var newContent = '';
      
      for (var i = 0; i < responseObject.length; i++) {    // Loop through object
        newContent += '<tr id="' + responseObject[i].Id + '">';
        newContent += '<td name="cstmname">' + responseObject[i].CstmName + '</td>';
        newContent += '<td name="cstmtype">' + responseObject[i].CstmType + '</td>';
        newContent += '<td name="city">' + responseObject[i].City + '</td>';
        newContent += '<td name="area">' + responseObject[i].Area + '</td>';
        newContent += '<td name="address">' + responseObject[i].Address + '</td>';
        newContent += '<td name="owner_name">' + responseObject[i].OwnerName + '</td>';
        newContent += '<td name="telephone">' + responseObject[i].Telephone + '</td>';
        newContent += '<td name="police">' + responseObject[i].Police + '</td>';
        newContent += '<td name="axis">' + responseObject[i].Axis + '</td>';
        newContent += '<td name="remark">' + responseObject[i].Remark + '</td>';
        newContent += '<td>' + '<input type="button" class="chg" onclick="upt(' +"'"+ responseObject[i].Id +"'"+ ')" value="修改" /> <input type="button" class="del" onclick="del(' +"'"+ responseObject[i].Id +"'"+ ')" value="删除" />' + '</td>';
        newContent += '</tr>';
      }
      // Update the page with the new content
      document.getElementById('rows').innerHTML = newContent;
      $('#customer_results').show()
      $('#export').show()
    }
  };
  xhr.open("POST", "/handle_customers", true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  xhr.send(postStr);
  console.log(postStr);
};
$('#customer_results').hide();
$('#export').hide();
$('#export').click(function(){
  table2xlsx('xlsx','customer_results');
});
let cstmSelects = getCustomerName();
$("#cstmname").autocomplete({
  source: cstmSelects
});

var el = document.getElementById('customer_query_some');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    customer_query(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    customer_query(e);
  });
}

var el2 = document.getElementById('customer_query_all');
if (el2.addEventListener) {
  el2.addEventListener('click', function (e) {
    customer_query(e);
  }, false);
} else {
  el2.attachEvent('onclick', function (e) {
    customer_query(e);
  });
}

function upt(cstmid) {
  let domStr='<legend><span>修改记录</span></legend>';
  $('#' + cstmid).children().each(function (index, element) {
    if($(element).attr("name") != undefined){
    domStr += '<label>'+index +'：</label><input id="update_' + $(element).attr("name") + '" value="' + $(element).text() + '">'
    }
  });
  domStr += '<input type="button" class="sv" onclick="sv(' +"'"+ cstmid +"'"+ ')" value="保存" /><input type="button" class="sv" onclick="hd()" value="取消" />';
  $("#change_table").html(domStr);
  $("#change_table").show();
  $("html,body").animate({scrollTop:$("#change_table").offset().top},1000);
}

function sv(cstmid) {
  $.post("handle_customers",{
  operation:'update',
  id:cstmid,
  cstmname : document.getElementById("update_cstmname").value.replace(/\s+/g,""),
  cstmtype : document.getElementById("update_cstmtype").value.replace(/\s+/g,""),
  gaddr : document.getElementById("update_city").value.replace(/\s+/g,""),
  gname : document.getElementById("update_area").value.replace(/\s+/g,""),
  gphone : document.getElementById("update_address").value.replace(/\s+/g,""),
  ivaddr : document.getElementById("update_owner_name").value.replace(/\s+/g,""),
  ivname : document.getElementById("update_telephone").value.replace(/\s+/g,""),
  ivphone : document.getElementById("update_police").value.replace(/\s+/g,""),
  ivphone : document.getElementById("update_axis").value.replace(/\s+/g,""),
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

function del(cstmid) {
  var cfm = confirm("确认要删这条记录吗？");
  if (cfm == true) {
    $.post("handle_customers",
      {
        operation: "delete",
        id: cstmid
      },
      function (data, status) {
        if (status == "success") {
          alert(data);
          $("#" +cstmid).hide();
        } else {
          alert("服务器错误，请联系周京成");
        }
      });
  }
  else {
    return;
  }
}
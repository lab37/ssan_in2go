function queryContract(e) {
  
  let tgt = getTarget(e);
  var postStr;
  if (tgt.id == 'contract_query_all') {
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
    postStr = postStr + '&ccsn=' + encodeURIComponent(document.getElementById("ccsn").value.replace(/\s+/g, ""));
    
    postStr = postStr + '&vector=' + encodeURIComponent(document.getElementById("vector").value.replace(/\s+/g, ""));
    
    postStr = postStr + '&create_date=' + encodeURIComponent(document.getElementById("create_date").value.replace(/\s+/g, ""));
    postStr = postStr + '&cstmname=' + encodeURIComponent(document.getElementById("cstmname").value.replace(/\s+/g, ""));
    postStr = postStr + '&prdtname=' + encodeURIComponent(document.getElementById("prdtname").value.replace(/\s+/g, ""));
  }

  var xhr = new XMLHttpRequest();

  xhr.onload = function () {
    if (xhr.status == 200) {
      console.log(xhr.responseText);
      responseObject = JSON.parse(xhr.responseText);
      var newContent = '';
      for (var i = 0; i < responseObject.length; i++) {
        newContent += '<tr id="' + responseObject[i].Id + '">';
        newContent += '<td name="id">' + responseObject[i].Id + '</td>';
        newContent += '<td name="ccsn">' + responseObject[i].Ccsn + '</td>';
        newContent += '<td name="vector">' + (responseObject[i].Vector==0?"购进":"售出") + '</td>';
        
        newContent += '<td name="create_date">' + responseObject[i].CreateDate + '</td>';
        newContent += '<td name="cctype">' + responseObject[i].CcType + '</td>';
        newContent += '<td name="ivtype">' + responseObject[i].IvType + '</td>';
        newContent += '<td name="cstmname">' + responseObject[i].CstmName + '</td>';
        newContent += '<td name="prdtname">' + responseObject[i].PrdtName + '</td>';
        newContent += '<td name="specific">' + responseObject[i].Specific + '</td>';
        newContent += '<td name="price">' + responseObject[i].Price + '</td>';
        newContent += '<td name="quantity">' + responseObject[i].Quantity + '</td>';
        newContent += '<td name="remark">' + responseObject[i].Remark + '</td>';
        newContent += '<td>' + '<input type="button" class="chg" onclick="upt(' +"'" + responseObject[i].Id +"'" +')" value="修改" /> <input type="button" class="del" onclick="del(' +"'"+ responseObject[i].Id +"'"+ ')" value="删除" />' + '</td>';
        newContent += '</tr>';
      }
      // Update the page with the new content
      document.getElementById('rows').innerHTML = newContent;
      $('#export').show();
      $("#contract_results").show();

    }
  };

  xhr.open("POST", "/handle_contracts", true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  console.log(postStr);
  xhr.send(postStr);
};

$("#contract_results").hide();

$('#export').hide();
$('#export').click(function(){
  table2xlsx('xlsx','contract_results');
});
$("#create_date").datepicker();

let cstmSelects = getCustomerName();
$("#cstmname").autocomplete({
  source: cstmSelects
});


var el = document.getElementById('contract_query_some');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    queryContract(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    queryContract(e);
  });
}

var el2 = document.getElementById('contract_query_all');
if (el2.addEventListener) {
  el2.addEventListener('click', function (e) {
    queryContract(e);
  }, false);
} else {
  el2.attachEvent('onclick', function (e) {
    queryContract(e);
  });
}


function upt(id) {
  let domStr='<legend><span>修改记录</span></legend>';
  $('#' + id).children().each(function (index, element) {
    if($(element).attr("name") != undefined){
    domStr += '<label>'+index +'：</label><input id="update_' + $(element).attr("name") + '" value="' + $(element).text() + '">'
    }
  });
  domStr += '<input type="button" class="sv" onclick="sv(' +"'"+ id+"'" + ')" value="保存" /><input type="button" class="sv" onclick="hd()" value="取消" />';
  $("#change_table").html(domStr);
  $("#change_table").show();
  $("html,body").animate({scrollTop:$("#change_table").offset().top},1000);
}

function sv(id) {
  $.post("handle_contracts", {
    operation: 'update',
    id: id,
    ccsn: document.getElementById("update_ccsn").value.replace(/\s+/g, ""),
    create_date: document.getElementById("update_create_date").value.replace(/\s+/g, ""),
    cctype: document.getElementById("update_cctype").value.replace(/\s+/g, ""),
    ivtype: document.getElementById("update_ivtype").value.replace(/\s+/g, ""),
    vector: document.getElementById("update_vector").value.replace(/\s+/g, "")=="购进"?"0":"1",
    cstmname: document.getElementById("update_cstmname").value.replace(/\s+/g, ""),
    prdtname: document.getElementById("update_prdtname").value.replace(/\s+/g, ""),
    specific: document.getElementById("update_specific").value.replace(/\s+/g, ""),
    price: document.getElementById("update_price").value.replace(/\s+/g, ""),
    quantity: document.getElementById("update_quantity").value.replace(/\s+/g, ""),
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
    $.post("handle_contracts",
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
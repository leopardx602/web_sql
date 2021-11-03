$(document).ready(() => {
  var table = $('#example').DataTable();

  $.get("/products", function(data) {
    Object.keys(data).forEach(key => {
      console.log(key, data[key]);
      table.row.add([
        data[key]['id'],
        data[key]['name'],
        data[key]['price'],
        data[key]['image'],
        data[key]['createAt'],
        data[key]['updateAt'],
      ]).draw( false );
    });
  });

  $('#example tbody').on( 'click', 'tr', function () {
    if ( $(this).hasClass('selected') ) {
        $(this).removeClass('selected');
    }
    else {
        table.$('tr.selected').removeClass('selected');
        $(this).addClass('selected');
    }
  });

  $("#add").click(()=>{
    console.log($("#addName").val())
    console.log($("#addPrice").val())

    js = JSON.stringify({
      "name":$("#addName").val(),
      "price":parseInt($("#addPrice").val(), 10)
    })
    console.log(js)

    $.post("/products", js, function(data, status){
      console.log(data, status)
    })
  })

  $("#delete").click(()=>{
    //table.row('.selected').remove().draw( false );
    //console.log(table.row(".selected").data()[0])
    console.log($("#deleteID").val())
    $.ajax({
      url: `/products/${$("#deleteID").val()}`,
      type: 'DELETE',
      success: function(result) {
        console.log(result)
      }
    });
  })

  $("#update").click(()=>{
    console.log($("#updateID").val())

    js = JSON.stringify({
      "id":parseInt($("#updateID").val(), 10),
      "name":$("#updateName").val(),
      "price":parseInt($("#updatePrice").val(), 10)
    })
    $.ajax({
      url: `/products/${$("#updateID").val()}`,
      type: 'put',
      contentType: "application/json",
      data:js,
      success: function(result) {
        console.log(result)
      }
    });
  })

})
"use strict";

$(document).ready(function () {
  $.ajaxSetup({
    beforeSend: function (xhr, settings) {
      xhr.setRequestHeader("X-CSRFToken", $("input[name=csrfmiddlewaretoken]").val());
    }
  });

  $("#form button").on("click", function (event) {
    var post_id = $("#content_bg").attr("post-id");
    var url = "/api/posts/"+post_id;
    var button = event.target;

    if (button.id == "update") {
      var payload = {
        // csrfmiddlewaretoken: $("input[name=csrfmiddlewaretoken]").val(),
        content: $("textarea").first().val(),
        exist: document.getElementsByName("exist")[0].checked
      }
      $.ajax({
        url: url,
        type: 'PUT',
        data: JSON.stringify(payload),
        headers: {
          "Content-Type": "application/json",
          // "X-CSRFToken": $("input[name=csrfmiddlewaretoken]").val()
        },
        success: function(data) {
          $("#update").addClass("green");
          setTimeout(function () {
            $("#update").removeClass("green");
          }, 1000);
        }
      })
    }
    else if (button.id == "delete") {
      $.ajax({
        url: url,
        type: "DELETE",
        beforeSend: function (xhr) {
          xhr.setRequestHeader("X-CSRFToken", $("input[name=csrfmiddlewaretoken]").val());
        },
        success: function (data) {
          console.log(data);
          $("#delete").addClass("green");
          setTimeout(function () {
            $("#delete").removeClass("green");
          }, 1000);
        }
      });
    }
  });
  $("#comment button").on("click", function (event) {
    var post_id = $("#content_bg").attr("post-id");
    var url = "/api/comments/"+post_id;
    console.log(url);
    $.post(url, {
              csrfmiddlewaretoken: $("input[name=csrfmiddlewaretoken]").val(),
              author: $("input[name=author]").val(),
              content: $("input[name=content]").val()
            }, 
          function (data, success) {
            if (data.success == true) {
              $("#comment button").addClass("green");
              setTimeout(function () {
                $("#comment button").removeClass("green");
              }, 1000);
            }
          });
  })
});

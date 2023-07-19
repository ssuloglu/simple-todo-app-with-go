

function changeSpanValue(name, increase=1) {
  spanName = document.getElementsByName(name);
  if (increase==1)
    spanName[0].textContent = (parseInt(spanName[0].textContent)+1).toString();
  else 
    spanName[0].textContent = (parseInt(spanName[0].textContent)-1).toString();
}

function changeBadge(el,info=1){
  if (info==1)
    el.next(".label-text").next(".badge")[0].className = "badge badge-info";
  else
    el.next(".label-text").next(".badge")[0].className = "badge badge-warning";
  }

$(document).ready(function() {
    $('input[type=checkbox]').change(function() {
     
      if (this.checked) {
        $(this).next(".label-text").css("text-decoration-line", "line-through");
        
        changeSpanValue("doneItems");
        changeSpanValue("waitingItems", 0);
        changeBadge($(this));
      } else {
        $(this).next(".label-text").css("text-decoration-line", "none");

        changeSpanValue("doneItems", 0);
        changeSpanValue("waitingItems");
        changeBadge($(this), 0);
      }
    });
  });
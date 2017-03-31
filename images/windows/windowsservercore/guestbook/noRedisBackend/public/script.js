$(document).ready(function() {
  var headerTitleElement = $("#header h1");
  var formElement = $("#envy-form");
  var submitElement = $("#envy-submit");
  var hostAddressElement = $("#guestbook-host-address");

  $.getJSON("env", function(result){
	$.each(result, function(i, field){
		var theDiv = document.getElementById("environmentDetails");
		theDiv.innerHTML = theDiv.innerHTML + "<li>" + i + " = " + field;
		//var content = document.createTextNode();
		//theDiv.appendChild(content);            
	});
  });

  var donothing = function(data) {
  }

  var handleSubmission = function(e) {
    e.preventDefault();	
    $.getJSON("rpush/envy/crashme", donothing);
    return false;
  }

  // colors = purple, blue, red, green, yellow
  var colors = ["#549", "#18d", "#d31", "#2a4", "#db1"];
  var randomColor = colors[Math.floor(5 * Math.random())];
  (function setElementsColor(color) {
    headerTitleElement.css("color", color);
    submitElement.css("background-color", color);
  })(randomColor);

  submitElement.click(handleSubmission);
  formElement.submit(handleSubmission);
  hostAddressElement.append(document.URL);
});

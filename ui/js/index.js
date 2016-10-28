$(document).ready(function() {
    init();
    console.log("init");
});

function init() {
    loadItems();
    addBtn();
    run();
    modal();
}

var array;
var currentIndex;

function run() {
    $(document).ready(function() {
        $('#runButton').click(function() {
          // runFunction(array[0].name)
            if (document.getElementById("results") != null){
                document.getElementById("results").remove();
            }
            if (typeof(currentIndex) != "undefined") {
                runFunction(array[currentIndex].name)
            } else {
                var modal = document.getElementById('myModal');

                modal.style.display = "inline-block"
            }
        });
    });
}

function showPage() {
    document.getElementById("loader").style.display = "none";
    // document.getElementById("myDiv").style.display = "block";
}



// var values = [{
//     name: 'Jonny Strömberg',
//     born: 1986
// }, {
//     name: 'Jonas Arnklint',
//     born: 1985
// }, {
//     name: 'Martina Elm',
//     born: 1986
// }];


// userList.add({
//     name: "Gustaf Lindqvist",
//     born: 1983
// });


function loadItems() {
    $.ajax({
        url: "http://128.107.18.112:8888/function",
        type: 'GET',
        dataType: 'json',
        success: function(data) {
            console.log(data)
            for (var i = 0, len = data.length; i < len; i++) {
                // console.log(data[i]);
                // $('#functions').on('click', '', function() {
                //     btnAddToCart(data.name);
                // });
            }
            console.log(data);
            var options = {
                valueNames: ['name'],
                item: '<li><h3 class="name"></h3></li>'
            };
            var values = data
            array = data
            var functionList = new List('functions', options, values);

        },
        error: function(request, error) {
            alert("Error: " + JSON.stringify(request));
        }
    });
}

function loadFile(index) {
    var editor = ace.edit("editor");
    var decodedString
    editor.getSession().setMode("ace/mode/javascript");

    if (array[index].type == "FUNCTION_TYPE_BLOB") {
        console.log(array[index].name);
        decodedString = atob(array[index].sourceblob);
          editor.setValue(decodedString);
    } else {
        $.ajax({
            url: array[index].sourceurl,
            type: 'GET',
            success: function(data) {
                decodedString = atob(data.sourceblob);
                editor.setValue(decodedString);
                console.log(decodedString);
            },
            error: function(request, error) {
                alert("Error: " + JSON.stringify(request));
            }
        });
    }
}

function runFunction(fileName) {
    var cl = new CanvasLoader('canvasloader-container');
    cl.setColor('#1475b5'); // default is '#000000'
    cl.setShape('spiral'); // default is 'oval'
    cl.setDiameter(15);
    cl.show(); // Hidden by default
    var param1 = document.getElementById("param1").value
    var param2 = document.getElementById("param2").value

    console.log(param1, param2);
    $.ajax({
        url: "http://128.107.18.112:8888/function/" + fileName + "/run",
        type: 'POST',
        dataType: 'json',
        data: JSON.stringify({
            runparams: [param1, param2]
        }),
        success: function(data) {
            cl.hide(); // Hidden by default

            console.log(data)
            var json = JSON.stringify(data, null, 4)
                // $().appendTo('#results');
                console.log(json);
                output(syntaxHighlight(json));


        },
        error: function(request, error) {
            alert("Error: " + JSON.stringify(request));
            cl.hide(); // Hidden by default

        }
    });
}


function addBtn() {
    $('ul li h3').live('click', function() {
        var index = $(this).parent('li').index();
        console.log(array[index].name);
        loadFile(index);
        currentIndex = index
    });
}

function modal() {
    // Get the modal
    var modal = document.getElementById('myModal');

    // Get the button that opens the modal
    var btn = document.getElementById("runButton");

    // Get the <span> element that closes the modal
    var span = document.getElementsByClassName("close")[0];

    // When the user clicks on <span> (x), close the modal
    span.onclick = function() {
        modal.style.display = "none";
    }

    // When the user clicks anywhere outside of the modal, close it
    window.onclick = function(event) {
        if (event.target == modal) {
            modal.style.display = "none";
        }
    }
}

function output(inp) {
  var param2 = document.getElementById("run")
  var child = document.createElement('pre')
  child.setAttribute("id", "results");
  param2.appendChild(child).innerHTML = inp;

}

function syntaxHighlight(json) {
    json = json.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
    return json.replace(/("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?)/g, function (match) {
        var cls = 'number';
        if (/^"/.test(match)) {
            if (/:$/.test(match)) {
                cls = 'key';
            } else {
                cls = 'string';
            }
        } else if (/true|false/.test(match)) {
            cls = 'boolean';
        } else if (/null/.test(match)) {
            cls = 'null';
        }
        return '<span class="' + cls + '">' + match + '</span>';
    });
}

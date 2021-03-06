$(document).ready(function() {
    init();
    console.log("init");
});

function init() {
    loadItems();
    addBtn();
    addDeleteAction();
    run();
    update();
    newFile();
    modal();
}

var array;
var currentIndex;

var ip = "http://128.107.18.112:8888"

var options = {
    valueNames: ['name'],
    item: '<li><h3 class="name"></h3> <button class="delete">Delete</button></li>'
};
var functionList = new List('functions', options);

var nameField = $('#nameFunction')
var newFileName

function run() {
    $(document).ready(function() {
        $('#runButton').click(function() {
            // runFunction(array[0].name)
            if (document.getElementById("results") != null) {
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

function update() {
    $(document).ready(function() {
        $('#updateButton').click(function() {
            if (typeof(currentIndex) != "undefined") {
                updateFunction(array[currentIndex].name)
            } else {
                var modal = document.getElementById('myModal');
                modal.style.display = "inline-block"
            }
        });
    });
}

function newFile() {
    $(document).ready(function() {
        $('#initFunction').click(function() {
            if (nameField.val() != "") {
                initFunction()
                create()
            } else {
                alert("Please name file first.")
            }
        });
    });
}

function create() {
    // Set color
    var create = document.getElementById('createButton');
    create.style.color = "#fff"
    create.style.backgroundColor = "#28a8e0"
    $(document).ready(function() {
        $('#createButton').click(function() {
            createFunction()
            create.style.color = "lightGray"
            create.style.backgroundColor = "gray"
        });
    });
}

function showPage() {
    document.getElementById("loader").style.display = "none";
    // document.getElementById("myDiv").style.display = "block";
}


function loadItems() {
    $.ajax({
        url: ip + "/function",
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

            var values = data
            array = data
            functionList.clear()
            functionList = new List('functions', options, values);

        },
        error: function(request, error) {
            alert("Error: " + JSON.stringify(request));
        }
    });
}



function loadFile(index) {
    var editor = ace.edit("editor");
    var decodedString
    editor.getSession().setOption("useWorker", false);
    editor.getSession().setMode("ace/mode/javascript");
    editor.focus();


    if (array[index].type == "FUNCTION_TYPE_BLOB") {
        console.log(array[index].name);
        decodedString = atob(array[index].sourceblob);
        editor.setValue(decodedString, -1);
    } else {
        $.ajax({
            url: array[index].sourceurl,
            type: 'GET',
            success: function(data) {
                decodedString = atob(data.sourceblob);
                editor.setValue(decodedString, -1);
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
        url: ip + "/function/" + fileName + "/run",
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

function updateFunction(fileName) {
    var cl = new CanvasLoader('canvasloader-container2');
    cl.setColor('#1475b5'); // default is '#000000'
    cl.setShape('spiral'); // default is 'oval'
    cl.setDiameter(15);
    cl.show(); // Hidden by default

    console.log("UPDATE");
    // Update
    //curl -H "Content-Type: application/json" -s -X PUT localhost:8888/function -d '{"name":"tropo2", "version":"2.1"}'
    var milliseconds = (new Date).getTime();
    var string = milliseconds.toString();

    console.log(fileName, milliseconds);
    var stringToEncode = editor.getValue();
    console.log(stringToEncode);
    encodedString = btoa(stringToEncode);

    $.ajax({
        url: ip + "/function",
        type: 'PUT',
        dataType: 'json',
        data: JSON.stringify({
            name: fileName,
            version: string,
            sourceblob: encodedString
        }),
        success: function(data) {
            cl.hide(); // Hidden by default

            console.log(data)
            var json = JSON.stringify(data, null, 4)
                // $().appendTo('#results');
            console.log(json);
            alert('Success');
            loadItems()


        },
        error: function(request, error) {
            alert("Error: " + JSON.stringify(request));
            cl.hide(); // Hidden by default

        }
    });
}

function initFunction() {
    newFileName = nameField.val()
    nameField.val("")

    editor.insert("Get started on your new function here, hit create function when done");
    editor.gotoLine(0);
    editor.focus();
    editor.selectAll();
}

function createFunction() {

    var stringToEncode = editor.getValue();
    encodedString = btoa(stringToEncode);

    console.log(encodedString);

    name = newFileName.replace(/\\/g, '/');
    name = name.substring(name.lastIndexOf('/') + 1, name.lastIndexOf('.'));

    console.log(name);
    console.log(newFileName);

    $.ajax({
        url: ip + "/function",
        type: 'POST',
        dataType: 'json',
        data: JSON.stringify({
            cachedir: ".cache",
            name: name,
            namespace: "default",
            sourceblob: encodedString,
            sourcefile: "testsource/" + newFileName,
            type: "FUNCTION_TYPE_BLOB"
        }),
        success: function(data) {
            alert('Successfully created ' + newFileName);
            loadItems()
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

        $('ul li h3').css('color', 'black');
        $(this).css({
            "color": "gray"
        });
    });
}

function addDeleteAction() {
    $('ul li button').live('click', function() {
        var index = $(this).parent('li').index();
        var result = confirm("Are you sure you want to delete " + array[index].name + "?");
        if (result) {
            //Logic to delete the item
            deleteFunction(index)
        }

    });
}

function deleteFunction(index) {
    console.log(array[index].name);
    var functionToDelete = array[index].name
    $.ajax({
        url: ip + "/function/" + array[index].name,
        type: 'DELETE',
        success: function(data) {
            alert("Successfully deleted: " + functionToDelete);
            loadItems()
        },
        error: function(request, error) {
            alert("Error: " + JSON.stringify(request));
        }
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

    child.setAttribute("id", "results");
    child.setAttribute("table-layout", "fixed");
    child.style.setAttribute('width', '100px')
    child.style.setAttribute('word-break', 'break-all')



}

function syntaxHighlight(json) {
    json = json.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
    return json.replace(/("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?)/g, function(match) {
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

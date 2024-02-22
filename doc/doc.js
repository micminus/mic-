document.addEventListener("DOMContentLoaded", function() {
    let names = document.querySelectorAll(".name");

    let colors = ['red', 'green', 'blue'];
    let currentColorIndex = 0;

    setInterval(function() {
        names.forEach(function (name) {
            name.style.color = colors[currentColorIndex];
        });

        currentColorIndex = (currentColorIndex + 1) % colors.length;
    }, 1500);

    let tr = '<input type="checkbox" class="chek" disabled checked><label for="chek"></label>';
    for (i=0; i<100; i++) {
        document.getElementById("main").innerHTML = document.getElementById("main").innerHTML.replace("<c>", tr);
    }
});
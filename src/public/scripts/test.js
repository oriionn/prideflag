const buttons = document.querySelectorAll(".choices button");

function choose(e) {
    let id = e.target.getAttribute("data-choice");
    let url = new URL(window.location.href)

    let searchParams = url.searchParams;
    searchParams.set("a", id);

    window.location.href = url;
}

buttons.forEach(button => button.addEventListener("click", choose))

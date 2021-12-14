window.onload = () => {
    paginate('table');

    const categoryId = document.getElementById("categoryIdSearched").value;
    document.getElementById("categoryId").value = categoryId;
}
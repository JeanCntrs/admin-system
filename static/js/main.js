const paginate = tableId => {
  $(`#${tableId}`).DataTable();
}

const alert = (title = 'Success', text = 'Your data has been saved', icon = 'success') => {
  Swal.fire({
    title,
    text,
    icon,
    allowOutsideClick: false
  });
}

const confirmation = (title = 'Are you sure?', text = 'If you are sure confirm the action') => {
  return Swal.fire({
    title,
    text,
    icon: 'warning',
    showCancelButton: true,
    confirmButtonColor: '#3085d6',
    cancelButtonColor: '#d33',
    confirmButtonText: 'Yes, confirm it!',
    allowOutsideClick: false
  });
}

const getDataTable = (url, tableHeaders, fields, elementId, showBtnEdit = false, showBtnDelete = false, propertyName = '', tableId = 'table', isPopup = true, isChecked = false) => {
  let table = `<table data-property-name="${propertyName}" id="${tableId}" class="table">`;

  fetch(url)
    .then(response => response.json())
    .then(response => {
      table += '<thead class="table-dark">';
      table += '<tr>';

      if (isChecked) {
        table += '<th>';
        table += '</th>';
      }

      tableHeaders.forEach(header => {
        table += `<th>${header}</th>`;
      });

      if (showBtnEdit || showBtnDelete) {
        table += '<th>Actions</th>';
      }

      table += '</tr>';
      table += '</thead>';
      table += '<tbody>';

      response.forEach(element => {
        table += '<tr>';

        if (isChecked) {
          table += '<td>';
          table += `<input id="chk${element[propertyName]}" class="checkbox" type="checkbox" />`;
          table += '</td>';
        }

        fields.forEach(field => {
          table += '<td>';
          table += element[field]
          table += '</td>';
        });

        if (showBtnEdit || showBtnDelete) {
          table += '<td>';

          if (showBtnEdit) {
            table += `<a
                        class="btn btn-primary"

                        ${isPopup
                ? `data-bs-toggle="modal"
                            data-bs-target="#staticBackdrop"
                            onclick="openModal(${element[propertyName]}, '${tableId}'); getEntityById(${element[propertyName]})"`
                : `onclick="getEntityById(${element[propertyName]})"`
              }
                      >
                        <svg
                          xmlns="http://www.w3.org/2000/svg"
                          width="16"
                          height="16"
                          fill="currentColor"
                          class="bi bi-pencil-square"
                          viewBox="0 0 16 16"
                        >
                          <path
                            d="M15.502 1.94a.5.5 0 0 1 0 .706L14.459 3.69l-2-2L13.502.646a.5.5 0 0 1 .707 0l1.293 1.293zm-1.75 2.456-2-2L4.939 9.21a.5.5 0 0 0-.121.196l-.805 2.414a.25.25 0 0 0 .316.316l2.414-.805a.5.5 0 0 0 .196-.12l6.813-6.814z"
                          />
                          <path
                            fill-rule="evenodd"
                            d="M1 13.5A1.5 1.5 0 0 0 2.5 15h11a1.5 1.5 0 0 0 1.5-1.5v-6a.5.5 0 0 0-1 0v6a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5v-11a.5.5 0 0 1 .5-.5H9a.5.5 0 0 0 0-1H2.5A1.5 1.5 0 0 0 1 2.5v11z"
                          />
                        </svg>
                      </a>`
          }

          if (showBtnDelete) {
            table += ` <a
                        class="btn btn-danger"
                        onclick="showDeleteModalById(${element[propertyName]})"
                      >
                        <svg
                          xmlns="http://www.w3.org/2000/svg"
                          width="16"
                          height="16"
                          fill="currentColor"
                          class="bi bi-trash"
                          viewBox="0 0 16 16"
                        >
                          <path
                            d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"
                          />
                          <path
                            fill-rule="evenodd"
                            d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z"
                          />
                        </svg>
                      </a>`
          }

          table += '</td>';
        }

        table += '</tr>';
      });

      table += '</tbody>';
      table += '</table>';
      document.getElementById(elementId).innerHTML = table;
      paginate(`${tableId}`);
    })
}

const getDataSelect = (url, value, name, elementId) => {
  let options = '';

  fetch(url).then(response => response.json()).then(response => {
    options += '<option value="" selected>Choose</option>';

    response.forEach(element => {
      options += `<option value="${element[value]}">${element[name]}</option>`
    });

    document.getElementById(elementId).innerHTML = options;
  })
}

const openModal = (id, tableId = 'table') => {
  clearInputs();

  const propertyName = document.getElementById(`${tableId}`).getAttribute('data-property-name').replace('Id', '').toLowerCase();
  if (id == 0) {
    document.getElementById('txtModalTitle').innerHTML = `Add ${propertyName}`;
  } else {
    document.getElementById('txtModalTitle').innerHTML = `Edit ${propertyName}`;
  }
}

const clearInputs = () => {
  const controls = document.getElementsByClassName('form-control');

  for (let index = 0; index < controls.length; index++) {
    controls[index].value = '';
  }
}

const showDeleteModalById = (id) => {
  confirmation().then((result) => {
    if (result.isConfirmed) {
      deleteEntity(id);
    }
  })
}

const getSelectedCheckbox = () => {
  const checkboxList = document.getElementsByClassName('checkbox');
  let idsString = '';

  for (let index = 0; index < checkboxList.length; index++) {
    const checkbox = checkboxList[index];
    
    if (checkbox.checked === true) {
      idsString += `${checkbox.id.replace('chk', '')}*`
    }
  }

  if (idsString.length > 0) idsString = idsString.substring(0, idsString.length - 1);
  

  return idsString;
}
(function () {
  'use strict'
  window.addEventListener(
    'load',
    function () {
      // Fetch all the forms we want to apply custom Bootstrap validation styles to
      var forms = document.getElementsByClassName('needs-validation')
      // Loop over them and prevent submission
      Array.prototype.filter.call(forms, function (form) {
        form.addEventListener(
          'submit',
          function (event) {
            if (form.checkValidity() === false) {
              event.preventDefault()
              event.stopPropagation()
            }
            form.classList.add('was-validated')
          },
          false
        )
      })
    },
    false
  )
})()

const elem = document.getElementById('reservation-dates-modal')
const rangepicker = new DateRangePicker(elem, {
  // Change date format
  format: 'yyyy-mm-dd',
})

function notifyAlert(text, type) {
  notie.alert({
    type,
    text,
  })
}

function notifyModal(title, text, icon, confirmButtonText) {
  Swal.fire({
    title,
    text,
    icon,
    confirmButtonText,
  })
}

const myHTML = `
    <form id="check-availability-form" action="" method="post" novalidate class="needs-validation">
      <div class="form-row">
          <div class="col">
              <div class="form-row" id="reservation-dates-modal">
                  <div class="col">
                      <input disabled required class="form-control" type="text" name="start" id="start" placeholder="Arrival date" autocomplete="off">
                      </div>
                  <div class="col">
                      <input disabled required class="form-control" type="text" name="end" id="end" placeholder="Departure" autocomplete="off">
                  </div>
              </div>
          </div>
      </div>
    </form>
  `

// Testing by clicking a button
document.getElementById('colorButton').addEventListener('click', () => {
  notifyAlert("This is a message", "success")
  // notifyModal('Title', 'Hello', 'success', 'All good!')
  // attention.toast({ title: 'Hello world', icon: 'error' })
  // attention.error({ title: 'Oops!' })
  // attention.custom({ html: myHTML })
  // console.log(myModule()["success"]())
})

let attention = alertModule()

// myModule encloses other functions, toast, success...
// if we imagine toast or success were thousands of lines of code
// now we can access these using the var attention above
function alertModule() {
  const toast = (c) => {
    // c will be overwritten by the const {} below, i.e. if is c is not specified
    // we use the default values below
    const { title = '', icon = 'success', position = 'top-end' } = c

    const Toast = Swal.mixin({
      toast: true,
      // using the values above
      title,
      position,
      icon,
      showConfirmButton: false,
      timer: 3000,
      timerProgressBar: true,
      didOpen: (toast) => {
        toast.addEventListener('mouseenter', Swal.stopTimer)
        toast.addEventListener('mouseleave', Swal.resumeTimer)
      },
    })
    // Run
    Toast.fire({})
  }

  const success = (c) => {
    const { title = '', icon = 'success', text = '', footer = '' } = c
    Swal.fire({
      icon,
      title,
      text,
      footer,
    })
  }

  const error = (c) => {
    const { title = '', icon = 'error', text = '', footer = '' } = c
    Swal.fire({
      icon,
      title,
      text,
      footer,
    })
  }

  const custom = async (c) => {
    const { title = '', html = '' } = c
    const { value: formValues } = await Swal.fire({
      title,
      html,
      backdrop: false,
      focusConfirm: false,
      showCancelButton: true,
      // To open date picker
      willOpen: () => {
        const elem = document.getElementById('reservation-dates-modal')
        const rp = new DateRangePicker(elem, {
          format: 'yyyy-mm-dd',
          showOnFocus: true,
        })
      },
      preConfirm: () => {
        return [
          // Values from myHTML above
          document.getElementById('start').value,
          document.getElementById('end').value,
        ]
      },
      didOpen: () => {
        document.getElementById('start').removeAttribute('disabled')
        document.getElementById('end').removeAttribute('disabled')
      }
    })

    if (formValues) {
      Swal.fire(JSON.stringify(formValues))
    }
  }

  return {
    toast,
    success,
    error,
    custom,
  }
}

// Popup module
function popupModule() {}

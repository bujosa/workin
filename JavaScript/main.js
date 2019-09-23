// Get DOM Elements
const modal = document.querySelector("#my-modal");
const modalBtn = document.querySelector("#modal-btn");
const closeBtn = document.querySelector(".close");
const saveBtn = document.querySelector("#saveForm");
let horaIncidente, categoriaIncidente, fechaIncidente, contentString;

//Hide all markers
let hideAllInfoWindows = map => {
  markers.forEach(marker => {
    marker.infowindow.close(map, marker);
  });
  infowindows.forEach(e => {
    e.close();
  });
};

// Open
let openModal = () => (modal.style.display = "block");

// Close
let closeModal = () => (modal.style.display = "none");

// Close If Outside Click
let outsideClick = e => {
  if (e.target == modal) {
    modal.style.display = "none";
  }
};

// Events
modalBtn.addEventListener("click", openModal);
closeBtn.addEventListener("click", closeModal);
window.addEventListener("click", outsideClick);

// Prevent Form refresh

saveBtn.addEventListener("click", e => {
  e.preventDefault();
  categoriaIncidente = Array.from(
    document.querySelectorAll('[name="categoria"]')
  ).filter(e => e.checked)[0].value;
  horaIncidente = document.querySelector('[type="time"]').value;
  fechaIncidente = document.querySelector('[type="date"]').value;
  contentString = `
  <b>Hora del incidente: </b> ${horaIncidente} <br />
  <b>Categoria de incidente: </b> ${categoriaIncidente} <br />
  <b>Fecha del incidente: </b> ${fechaIncidente}
  `;
  closeModal();
  var infowindow = new google.maps.InfoWindow({
    content: contentString
  });

  map.addListener("click", function(e) {
    let res = placeMarker(e.latLng, map, infowindow);
    google.maps.event.addListener(res, "click", function() {
      hideAllInfoWindows(map);
      this.infowindow.open(map, this);
    });
    infowindows.push(infowindow);
    google.maps.event.clearInstanceListeners(map);

    let entradas = document.querySelectorAll("input");

    entradas.forEach(i => {
      if (entradas[i].type === "radio") entradas[i].checked = false;
      else {
        if (entradas[i].type !== "submit") entradas[i].value = "";
      }
    });
    document.querySelector("textarea").value = "";
  });
});

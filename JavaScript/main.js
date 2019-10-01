// Get DOM Elements
const modal = document.querySelector("#my-modal");
const modalBtn = document.querySelector("#modal-btn");
const closeBtn = document.querySelector(".close");
const saveBtn = document.querySelector("#saveForm");
let horaIncidente,
  categoriaIncidente,
  fechaIncidente,
  contentString,
  descripcion,
  modoDelincuente;

let postLocationData = async (datos = {}) => {
  response = await fetch("http://localhost:8080/LocationAdd", {
    method: "POST",
    mode: "no-cors",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(datos)
  });
};

let postCrimeData = async (datos = {}) => {
  response = await fetch("http://localhost:8080/CrimeAdd", {
    method: "POST",
    mode: "no-cors",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(datos)
  });
};

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
  descripcion = document.getElementById("textoDescripcion").value;
  // modoDelincuente = Array.from(
  //   document.querySelectorAll('[name="modo"]')
  // ).filter(e => e.checked)[0].value;
  modoDelincuente = "pasola";
  contentString = `
  <b>Hora del incidente: </b> ${horaIncidente} <br />
  <b>Categoria de incidente: </b> ${categoriaIncidente} <br />
  <b>Fecha del incidente: </b> ${fechaIncidente} <br />
  <b>Descripcion: </br> ${descripcion}
  `;
  let crimeData = {
    Id: 1,
    Loc: 123.354,
    Cat: categoriaIncidente,
    Date: fechaIncidente,
    Time: horaIncidente,
    Description: descripcion,
    Mode: modoDelincuente
  };

  closeModal();
  var infowindow = new google.maps.InfoWindow({
    content: contentString
  });

  map.addListener("click", function(e) {
    //console.log(`Latitud: ${e.latLng.lat()}, Longitud: ${e.latLng.lng()}`);
    postLocationData({
      User: 2,
      Latitude: e.latLng.lat(),
      Longitude: e.latLng.lng()
    });
    let res = placeMarker(e.latLng, map, infowindow);
    google.maps.event.addListener(res, "click", function() {
      hideAllInfoWindows(map);
      this.infowindow.open(map, this);
    });
    infowindows.push(infowindow);
    google.maps.event.clearInstanceListeners(map);

    let entradas = Array.from(document.querySelectorAll("input"));
    debugger;
    entradas.forEach(i => {
      if (i.type === "radio") i.checked = false;
      else {
        if (i.type !== "submit") i.value = "";
      }
    });
    document.querySelector("textarea").value = "";
    postCrimeData(crimeData);
  });
});

document.body.addEventListener("linkCreated", function (evt) {
  const data = evt.detail;  

  if (!data || !data.short) {
    console.warn("Data linkCreated tidak valid:", data);
    return;
  }

  const STORAGE_KEY = "u_short_history";
  let history = JSON.parse(localStorage.getItem(STORAGE_KEY) || "[]");

  const isDuplicate = history.some((item) => item.short === data.short);

  if (!isDuplicate) {
    history.unshift({
      id: data.id,
      original: data.original,
      short: data.short,
      date: new Date().toLocaleTimeString(),
    });

    localStorage.setItem(STORAGE_KEY, JSON.stringify(history.slice(0, 10)));

    console.log("Local Storage Updated: ", data.short);
  }
});
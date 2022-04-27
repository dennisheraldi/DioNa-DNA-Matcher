
export default function fetchGET(url) {
  var data = null;
  var isPending = true;
  var error = null;
  const abortCont = new AbortController();
  return fetch(url, { signal: abortCont.signal })
    .then(res => {
      if (res.status === 200) {
        return res.json();
      } else if (res.status === 400) {
        return res.json().then(data => {throw new Error(data.error)})
      } else {
        throw new Error('Something went wrong');
      }
    })
    .then(dat => {
      isPending = false;
      data = dat;
      error = null;
      return { data, isPending, error };
    })
    .catch(err => {
      if (err.name === 'AbortError') {
        console.log('fetch aborted')
      } else {
        // auto catches network / connection error
        isPending = false;
        error = err.message;
      }
      return { data, isPending, error };
    })
}
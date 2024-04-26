import { useState } from "react";
import Header from "./components/Header"
import Home from "./components/Home";
import SearchButton from "./components/SearchButton";
import SearchInput from "./components/SearchInput";
import Swal from "sweetalert2";
useState
function App() {
  const [linkValue, setLinkValue] = useState({
    startLink : '',
    endLink : '',
  })
  const [isError,setIsError] = useState([])
  const [resultResponse, setResultResponse] = useState({
    exec : 0,
    len : 0,
    urls : [],
    resultLink: [],
  })
  function handleInputSearch(value,type){
    if (type === "Start"){
      setLinkValue(prevState => {
        return {
          ...prevState,
          startLink: value
        }
      })
    } else if (type === "End"){
      setLinkValue(prevState => {
        return {
          ...prevState,
          endLink: value
        }
      })
    }
  }

  const errorMessage = () => {
    let showError = "Title "
    showError += isError[0]
    if (isError.length > 1) {
      showError += " and "
      showError += isError[1]
    }
    showError+= " dont exist!"
    Swal.fire({
      title: 'Error!',
      text: showError,
      icon: 'error',
      confirmButtonText: 'Try Another Title'
    })
    console.log(isError)
    setIsError([])
  }
  return (
    <div className="h-screen bg-[#68649c]">
      {isError.length !== 0 && errorMessage()}
      <main>
        <Header />
        <Home />
        <div className="gap-10">
          <div className="mt-10 flex flex-row gap-10 justify-center items-center">
            <SearchInput type="Start"  setLinkValue={handleInputSearch} />
            <h1 className="font-spaceComics text-[#f8ec34] text-4xl">TO</h1>
            <SearchInput type="End"  setLinkValue={handleInputSearch} />
          </div>
          <div className="mt-20 mb-10 flex justify-center items-center">
            <SearchButton 
              linkValue={linkValue} 
              isError={setIsError}
              setResultResponse={setResultResponse}/>
          </div>
        </div>

      </main>
    </div>
  );
}

export default App;
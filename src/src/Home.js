import React, { useState, useEffect } from 'react';
import './App.css';
import './Home.css';
import SearchButton from './components/SearchButton';
import OnOff from './components/OnOff';
import "./fonts/Inter-Bold.ttf"
import W from "./assets/w.png"
import Namespaces from './components/Namespaces';
import ResList from './components/ResList';
import SearchInput from './components/SearchInput'
import Greph from './components/Greph';
import Swal from 'sweetalert2'
function Home() {
  const [linkValue, setLinkValue] = useState({
    startLink : '',
    endLink : '',
  })
  const [isOn, setIsOn] = useState(false);
  const [isName, setIsName] = useState(false);
  const [isError, setIsError] = useState([]);
  const [resultResponse, setResultResponse] = useState({
    exec : 0,
    len : 0,
    urls : [],
    resultLink: [],
  })

  const [nodes, setNodes] = useState([]);
  const [links, setLinks] = useState([]);

  useEffect(() => {
    if (resultResponse.resultLink.length > 0) {
      setNodes(resultResponse.resultLink.map((name, index) => ({
        id : index+1,
        name : name
      })))
    }
  }, [resultResponse.resultLink]);

  useEffect(() => {
    if (resultResponse.resultLink.length > 1) {
      const newLinks = [];
      for (let i = 0; i < resultResponse.resultLink.length - 1; i++) {
        newLinks.push({ source: i + 1, target: i + 2 });
      }
      setLinks(newLinks);
    }
  }, [resultResponse.resultLink]);

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
    <div>
      {isError.length !== 0 && errorMessage()}
      <div className='topnav'>
        <img src={W} width='70px' height='70px' alt=''></img>
        <h1 style={{fontFamily:"Inter",color:"#ffffff",display:"inline-block"}}>WIKIBOWO DA WIKIRACER</h1>
      </div>
      <div className='first'>
        <p>"Balapan adalah jalan hidupku. Curang pake Website adalah cara menangku"</p>
        <p>- Sun Tzu, Art of War</p>
      </div>
      <div className='second'>
          <OnOff isOn={isOn} setIsOn={setIsOn}/>
          <Namespaces isName={isName} setIsName={setIsName}/>
        <div>
          <SearchInput setLinkValue= {handleInputSearch} type="Start"/>
          <SearchInput setLinkValue= {handleInputSearch} type="End"/>
        </div>
        <SearchButton 
          isOn={isOn} 
          linkValue={linkValue} 
          setResultResponse={setResultResponse} 
          isName={isName} 
          isError={setIsError}/>
      </div>
      <div>
        <ResList resultResponse={resultResponse}/>
      </div>
      {resultResponse.resultLink.length !== 0 &&
        <div className='fourth'>
          <Greph nodes={nodes} links={links} tempLinks={resultResponse.resultLink}></Greph>
        </div>
      }
    </div>
  );
}

export default Home;

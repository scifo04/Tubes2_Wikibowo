import React, { useState } from 'react'; 
import './App.css';
import './Home.css';
// import LinkForm from './components/LinkForm';
import SearchButton from './components/SearchButton';
// import EndForm from './components/EndForm';
import OnOff from './components/OnOff';
import "./fonts/Inter-Bold.ttf"
import W from "./assets/w.png"
import Namespaces from './components/Namespaces';
import ResList from './components/ResList';
import SearchInput from './components/SearchInput'

function Home() {
  const [linkValue, setLinkValue] = useState({
    startLink : '',
    endLink : '',
  })
  const [isOn, setIsOn] = useState(false);
  const [isName, setIsName] = useState(false);
  const [open, setOpen] = useState(''); //untuk open popup message
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
  return (
    <div>
      {open && console.log("Process")}
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
          setOpen={setOpen}/>
      </div>
      <div>
        <ResList resultResponse={resultResponse}/>
      </div>
    </div>
  );
}

export default Home;
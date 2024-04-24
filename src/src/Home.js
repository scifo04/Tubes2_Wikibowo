import React, { useState } from 'react'; 
import './App.css';
import './Home.css';
import LinkForm from './components/LinkForm';
import SearchButton from './components/SearchButton';
import EndForm from './components/EndForm';
import OnOff from './components/OnOff';
import "./fonts/Inter-Bold.ttf"
import W from "./assets/w.png"
import Namespaces from './components/Namespaces';
import ResList from './components/ResList';

function Home() {
  const [tempLinks, setTempLinks] = useState('');
  const [linkValue, setLinkValue] = useState('');
  const [finValue, setFinValue] = useState('');
  const [isOn, setIsOn] = useState(false);
  const [isName, setIsName] = useState(false);
  const [exec, setExec] = useState(0);
  const [len, setLen] = useState(0);
  const [urls, setUrls] = useState('');
  const [open, setOpen] = useState('');

  console.log(tempLinks);

  return (
    <div>
      {open && (
        <div className="popup-overlay">
          <div className="popup">
            <p style={{position:"absolute",top: "35%",left: "50%",transform: "translate(-50%, -50%)"}}>Please wait...</p>
          </div>
        </div>
      )}
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
          <LinkForm linkValue={finValue} setLinkValue={setLinkValue}/>
          <EndForm finValue={finValue} setFinValue={setFinValue}/>
        </div>
        <SearchButton isOn={isOn} linkValue={linkValue} finValue={finValue} isName={isName} tempLinks={tempLinks} setTempLinks={setTempLinks} setExec={setExec} setLen={setLen} setUrls={setUrls} setOpen={setOpen}/>
      </div>
      <div>
          <ResList  tempLinks={tempLinks} urls={urls} exec={exec} len={len}/>
      </div>
    </div>
  );
}

export default Home;
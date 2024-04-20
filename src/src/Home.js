import React, { useState } from 'react'; 
import './App.css';
import './Home.css';
import LinkForm from './components/LinkForm';
import SearchButton from './components/SearchButton';
import EndForm from './components/EndForm';
import OnOff from './components/OnOff';
import "./fonts/Inter-Bold.ttf"
import W from "./assets/w.png"

function Home() {
  const [tempLinks, setTempLinks] = useState('');
  const [linkValue, setLinkValue] = useState('');
  const [finValue, setFinValue] = useState('');
  const [isOn, setIsOn] = useState(false);

  console.log(tempLinks);

  return (
    <div>
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
        <div>
          <LinkForm linkValue={finValue} setLinkValue={setLinkValue}/>
          <EndForm finValue={finValue} setFinValue={setFinValue}/>
        </div>
        <SearchButton isOn={isOn} linkValue={linkValue} finValue={finValue} tempLinks={tempLinks} setTempLinks={setTempLinks}/>
      </div>
    </div>
  );
}

export default Home;
import logo from './logo.svg';
import './App.css';
import LinkForm from './components/LinkForm';
import SearchButton from './components/SearchButton';

function Home() {
  return (
    <div>
        <h1>WikiBowo Da WikiRacist</h1>
        <p>Balapan adalah jalan hidupku. Kapas adalah kerjaan budakku</p>
        <LinkForm/>
        <SearchButton/>
    </div>
  );
}

export default Home;
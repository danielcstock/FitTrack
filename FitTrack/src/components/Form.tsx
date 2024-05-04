
import {useState} from "react";
import axios from "axios";

export default function Form(){
    const [exercise, setExercise] = useState({name: "", times: "", weight: ""});

    function handleSubmit(e){ 
      const name = e.name;
      const times = e.times;
      const Weight = e.weight;

      axios.post("http://127.0.0.1:8080", {
        Name: name,
        TImes: times,
        weight: Weight
      })     
      e.preventDefault();
      console.log(exercise);
    }

    

    return (
        <>
          <form>
              <label>
                Exercício: <input name="name" onChange={(e) => setExercise({...exercise, name: e.target.value})} />
              </label><br></br>
    
              <label>
                Repetições: <input name="times" onChange={(e) => setExercise({...exercise, times: e.target.value})} />
              </label><br></br>
    
              <label>
                Peso: <input name="weight" onChange={(e) => setExercise({...exercise, weight: e.target.value})}/> kg
              </label><br></br>
    
              <button type="submit" onClick={(e) => handleSubmit(e)}>Enviar</button>
          </form>
        </>
      )
}
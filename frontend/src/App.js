import './App.css';
import Habit from "./components/Habit";
import HabitGroup from "./components/HabitGroup";
import HabitWindow from "./components/HabitWindow";

function App() {
  return (
    <div className="App" style={{
        height: "100vh"
    }}>
        <HabitWindow>
            <HabitGroup>
                <Habit></Habit>
                <Habit></Habit>
                <Habit></Habit>
            </HabitGroup>
            <HabitGroup>
                <Habit></Habit>
                <Habit></Habit>
                <Habit></Habit>
                <Habit></Habit>
            </HabitGroup>
            <HabitGroup>
                <Habit></Habit>
                <Habit></Habit>
                <Habit></Habit>
                <Habit></Habit>
                <Habit></Habit>
            </HabitGroup>
        </HabitWindow>
    </div>
  );
}

export default App;

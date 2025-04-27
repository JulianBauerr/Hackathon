function HabitWindow({children}) {
    return (
        <div style={{
            backgroundColor: "#D9D9D9",
            width: "55%",
            height: "70%",
            borderRadius: "40px",
            paddingTop: "20px",
            paddingLeft: "20px",

            overflowX: "scroll",
            gap: "10px",
            display: "flex",
            flexDirection: "row",

        }}>
            <div>
                {children}
            </div>
        </div>
    )
}

export default HabitWindow;
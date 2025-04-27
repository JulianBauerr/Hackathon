function HabitGroup({children}) {
    return (
        <div style={{
            backgroundColor: "#E7E7E7",
            width: "40%",
            height: "calc(100% - 65px)",
            display: "flex",
            borderRadius: "20px",
            paddingTop: "45px",

            flexDirection: "column",
            alignItems: "center",
            gap: "5px",
            flexShrink: "0",
        }}>
            {children}
        </div>
    )
}

export default HabitGroup
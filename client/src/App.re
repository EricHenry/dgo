[@react.component]
let make = () =>
    <>
        <MaterialUi_CssBaseline />
        <Navbar />
        MaterialUi.(
            <Grid container=true>
                <Grid item={true} xs=`V2>
                    "Grid item"->React.string
                </Grid>
                <Grid item={true} xs={`V3}>
                    "Another Grid item"->React.string 
                </Grid>
            </Grid>
        )
    </>
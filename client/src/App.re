let gameDescription = "This is a description of the game";
[@react.component]
let make = () => 
    <>
        <MaterialUi_CssBaseline />
        <Navbar />
        MaterialUi.(
            <Box component={`String("span")}>
                <Container>
                    <Grid container=true alignItems=`Center direction=`Column>
                        <Grid item=true>
                            <Typography variant=`H4>gameDescription</Typography>
                        </Grid>
                        <Grid item=true>
                            "Another Grid item"->React.string 
                        </Grid>
                    </Grid>
                </Container>
            </Box>
        )
    </>

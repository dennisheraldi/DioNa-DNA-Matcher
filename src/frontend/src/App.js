import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'
import Checkup from './pages/Checkup'
import { createTheme, ThemeProvider } from '@mui/material/styles'
import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles/';
import { purple } from '@material-ui/core/colors'
import Layout from './components/Layout'
import AddPenyakit from './pages/AddPenyakit'
import CheckHistory from './pages/CheckHistory'
const { palette } = createTheme()
const cust_theme = createTheme({
  palette: {
    primary: {
      main: '#6697BE'
    },
    secondary: {
      main: '#CE5B5F'
    },
  },
  typography: {
    fontFamily: 'Quicksand',
    fontWeightLight: 400,
    fontWeightRegular: 500,
    fontWeightMedium: 600,
    fontWeightBold: 700,
  }
})

const typo_theme = createMuiTheme({
  typography: {
    fontFamily: 'Quicksand',
    fontWeightLight: 400,
    fontWeightRegular: 500,
    fontWeightMedium: 600,
    fontWeightBold: 700,
  }
})

function App() {
  return (

      <ThemeProvider theme={cust_theme}>
        <MuiThemeProvider theme={typo_theme}>
        <Router>
          <Layout>
            <Switch>
              <Route exact path="/">
                <Checkup />
              </Route>
              <Route path="/add-penyakit">
                <AddPenyakit />
              </Route>
              <Route path="/check-history">
                <CheckHistory />
              </Route>
            </Switch>
          </Layout>
        </Router>
        </MuiThemeProvider>
      </ThemeProvider>
  );
}

export default App;

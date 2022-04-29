import React, { useState, useEffect } from "react";
import TextField from "@mui/material/TextField";
import Typography from "@material-ui/core/Typography";
import Box from "@mui/material/Box";
import Grid from "@mui/material/Grid";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import { Button } from "@material-ui/core";
import fetchGET from "../components/fetchGET";
import { Paper } from "@mui/material";

export default function CheckHistory() {
  const pageTitle = "Cek Riwayat - DioNA";
  useEffect(() => {
    document.title = pageTitle;
  }, [pageTitle]);
  const endpoint = `http://localhost:8080/v1/riwayat`;
  const [resDisplay, setResDisplay] = useState(false);
  const [data, setData] = useState(null);
  const [isPending, setIsPending] = useState(false);
  const [error, setError] = useState(null);
  /* 
    tanggal, nama, penyakit, kemiripan, diagnosa
    */
  const reF = /^\d{4}-\d{2}-\d{2}$/;
  const reC = /^\d{4}-\d{2}-\d{2}/;
  const reAll = /^[aA][lL][lL]$/;
  const reWS = /^\s*$/;
  const [kueri, setKueri] = useState("");
  const [link, setLink] = useState(null);

  useEffect(() => {
    const doAsync = async () => {
      if (link) {
        setData(null);
        setIsPending(true);
        setError(null);
        const { data: d, isPending: p, error: err } = await fetchGET(link);
        setIsPending(p);
        if (err) {
          setError(err);
        } else {
          setData(d);
        }
      }
    };
    doAsync();
  }, [link]);

  const HandleSubmit = (e) => {
    e.preventDefault();
    setResDisplay(true);
    setLink(`${endpoint}?query=${kueri.trim()}`);
  };
  return (
    <Card>
      <CardContent>
        <Grid sx={{ flexGrow: 1, p: 3 }} container columnSpacing={2}>
          <Box component={Grid} item align="center" boxShadow={0} xs={12} sx={{ mb: 5, fontSize: 16 }} fontWeight="fontWeightBold">
            <Typography component="div">
              <Box sx={{ fontSize: "h4.fontSize", fontWeight: "bold" }}>Cari Riwayat</Box>
            </Typography>
          </Box>
          <Box component={Grid} item boxShadow={0} xs={12}>
            <Box component="form" onSubmit={HandleSubmit}>
              <Grid sx={{ flexGrow: 1 }} container columnSpacing={2}>
                <Box component={Grid} item boxShadow={0} xs={3} sx={{ display: { xs: "none", md: "block" } }}></Box>
                <Box component={Grid} item boxShadow={0} xs={12} md={6}>
                  <TextField required fullWidth id="inp_kueri" label="Kueri Pencarian" onChange={(e) => setKueri(e.target.value)} />
                  <Typography component="div">
                    <Box
                      sx={{
                        color: "secondary.main",
                        fontWeight: "bold",
                        fontSize: 14,
                        pl: 1,
                        mt: 1,
                      }}
                    >
                      contoh: '2020-10-20', 'HIV', '2020-10-20 HIV', atau 'All'
                    </Box>
                  </Typography>
                </Box>
                <Box component={Grid} item boxShadow={0} xs={3} sx={{ display: { xs: "none", md: "block" } }}></Box>
                <Box component={Grid} item boxShadow={0} xs={3} sx={{ display: { xs: "none", md: "block" } }}></Box>
                <Box component={Grid} align="right" item boxShadow={0} xs={12} md={6}>
                  <Button type="submit" color="secondary" variant="contained">
                    SEARCH
                  </Button>
                </Box>

                <Box component={Grid} item boxShadow={0} xs={3} sx={{ display: { xs: "none", md: "block" } }}></Box>
              </Grid>
            </Box>
          </Box>
          <Box component={Grid} item xs={12} sx={{ display: resDisplay ? "block" : "none" }}>
            <Grid container columnSpacing={2}>
              <Box component={Grid} item align="center" boxShadow={0} xs={12} sx={{ mb: 3 }}>
                <Typography component="div">
                  <Box sx={{ fontSize: "h5.fontSize", fontWeight: "bold" }}>Result</Box>
                </Typography>
              </Box>
              <Box component={Grid} item align="center" boxShadow={0} xs={12}>
                <Typography component="div">
                  <Box sx={{ fontSize: "h5.fontSize", fontWeight: "bold", display: isPending ? "block" : "none" }}>Loading...</Box>
                  <Box sx={{ color: "red", fontSize: "h5.fontSize", fontWeight: "bold", display: error ? "block" : "none" }}>Error: {error}</Box>
                </Typography>
              </Box>
              <Box component={Grid} item boxShadow={0} xs={3} sx={{ mb: 5, display: { xs: "none", md: "block" } }}></Box>
              <Box component={Grid} item align="left" boxShadow={0} xs={12} lg={6} style={{ display: data ? "block" : "none", maxHeight: 200, overflow: "auto" }}>
                {data &&
                  data.data.map((item, index) => (
                    <Paper elevation={3} sx={{ mb: 3 }} key={index}>
                      <Card>
                        <CardContent>
                          <Typography component="div">
                            <Box sx={{ fontSize: { xs: 16, md: "h6.fontSize" }, fontWeight: "bold" }}>
                              {item.tanggal_pred} - {item.nama_pasien} - {item.nama_penyakit} - {item.status} - {item.similarity}%
                            </Box>
                          </Typography>
                        </CardContent>
                      </Card>
                    </Paper>
                  ))}
              </Box>
              <Box component={Grid} item boxShadow={0} xs={3} sx={{ mb: 5, display: { xs: "none", md: "block" } }}></Box>
            </Grid>
          </Box>
        </Grid>
      </CardContent>
    </Card>
  );
}

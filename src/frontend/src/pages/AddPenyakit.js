import React, { useState, useEffect } from "react";
import TextField from "@mui/material/TextField";
import Typography from "@material-ui/core/Typography";
import Box from "@mui/material/Box";
import Grid from "@mui/material/Grid";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import { Button } from "@material-ui/core";
import { Stack } from "@mui/material";
import { green } from "@mui/material/colors";
import fetchPOST from "../components/fetchPOST";

export default function AddPenyakit() {
  const pageTitle = "Tambah Penyakit - DioNA";
  useEffect(() => {
    document.title = pageTitle;
  }, [pageTitle]);
  const endpoint = "http://localhost:8080/v1/penyakit";
  const [resDisplay, setResDisplay] = useState(false);
  /* 
    tanggal, nama, penyakit, kemiripan, diagnosa
    */
  const [data, setData] = useState(null);
  const [isPending, setIsPending] = useState(false);
  const [error, setError] = useState(null);
  const fRead = new FileReader();
  const [penyakit_input, setPenyakitInput] = useState("");
  const [dna_input, setDnaInput] = useState("");
  const [payload, setPayload] = useState(null);
  const [dna_filename, setDnaFileName] = useState("No File Selected");
  const sxsuc = {
    color: green[500],
    fontWeight: "bold",
    fontSize: "h6.fontSize",
  };
  const sxerr = {
    color: "red",
    fontWeight: "bold",
    fontSize: "h6.fontSize",
  };
  useEffect(() => {
    const doAsync = async () => {
      if (payload) {
        setData(null);
        setIsPending(true);
        setError(null);
        const { data: d, isPending: p, error: err } = await fetchPOST(endpoint, payload);
        setIsPending(p);
        if (err) {
          setError(err);
        } else {
          setData(d);
        }
      }
    };
    doAsync();
  }, [payload]);
  const fileHandler = (e) => {
    fRead.onload = async (ev) => {
      const tdat = ev.target.result;
      setDnaInput(tdat);
    };
    fRead.readAsText(e.target.files[0]);
    setDnaFileName(e.target.files[0].name);
  };
  const handleSubmit = (e) => {
    e.preventDefault();
    setResDisplay(true);
    setPayload({
      nama_penyakit: penyakit_input.trim(),
      dna_penyakit: dna_input.trim(),
    });
  };
  return (
    <Card>
      <CardContent>
        <Grid sx={{ flexGrow: 1, p: 3 }} container columnSpacing={2}>
          <Box component={Grid} item align="center" boxShadow={0} xs={12} sx={{ mb: 5, fontSize: 16 }} fontWeight="fontWeightBold">
            <Typography component="div">
              <Box sx={{ fontSize: "h4.fontSize", fontWeight: "bold" }}>Tambah Penyakit</Box>
            </Typography>
          </Box>
          <Box component={Grid} item boxShadow={0} xs={12}>
            <Box component="form" onSubmit={handleSubmit}>
              <Grid sx={{ flexGrow: 1 }} container columnSpacing={2}>
                <Box component={Grid} item boxShadow={0} xs={2} sx={{ mb: 5, display: { xs: "none", md: "block" } }}></Box>
                <Box component={Grid} item boxShadow={0} xs={12} md={4} sx={{ mb: 5 }}>
                  <TextField required fullWidth id="inp_penyakit" label="Nama Penyakit" onChange={(e) => setPenyakitInput(e.target.value)} />
                </Box>
                <Box component={Grid} item boxShadow={0} xs={12} md={4} sx={{ mb: 5 }}>
                  <Box sx={{ fontSize: 12 }}>DNA Sequence (.txt, max 255 char): {dna_filename}</Box>
                  <Box component={Stack} direction="column" justifyContent="center">
                    <Button variant="contained" component="label">
                      Upload File
                      <input required onChange={fileHandler} type="file" accept=".txt" hidden />
                    </Button>
                  </Box>
                </Box>
                <Box component={Grid} item boxShadow={0} xs={2} sx={{ mb: 5, display: { xs: "none", md: "block" } }}></Box>
                <Box component={Grid} item boxShadow={0} xs={6} sx={{ mb: 5, display: { xs: "none", md: "block" } }}></Box>
                <Box component={Grid} align="right" item boxShadow={0} xs={12} md={4} sx={{ mb: 5 }}>
                  <Button type="submit" color="secondary" variant="contained">
                    ADD
                  </Button>
                </Box>
                <Box component={Grid} item boxShadow={0} xs={2} sx={{ mb: 5, display: { xs: "none", md: "block" } }}></Box>
              </Grid>
            </Box>
          </Box>
          <Box component={Grid} align="center" item xs={12} sx={{ display: resDisplay ? "block" : "none" }}>
            <Typography component="div">
              <Box sx={{ fontSize: "h5.fontSize", fontWeight: "bold", display: isPending ? "block" : "none" }}>Loading...</Box>
              <Box sx={[{ display: error ? "block" : "none" }, sxerr]}>Error: {error}</Box>
              <Box sx={[{ display: data ? "block" : "none" }, sxsuc]}>Success Adding {penyakit_input}</Box>
            </Typography>
          </Box>
        </Grid>
      </CardContent>
    </Card>
  );
}

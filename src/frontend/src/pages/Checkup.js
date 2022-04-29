import React, { useState, useEffect } from "react";
import TextField from "@mui/material/TextField";
import Typography from "@material-ui/core/Typography";
import Box from "@mui/material/Box";
import Grid from "@mui/material/Grid";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import { Button } from "@material-ui/core";
import { Stack } from "@mui/material";
import fetchPOST from "../components/fetchPOST";
import { green } from "@mui/material/colors";

export default function Checkup() {
  const pageTitle = "Cek Penyakit - DioNA";
  useEffect(() => {
    document.title = pageTitle;
  }, [pageTitle]);
  const endpoint = "http://localhost:8080/v1/riwayat";
  const [resDisplay, setResDisplay] = useState(false);
  /* 
    tanggal, nama, penyakit, kemiripan, diagnosa
    */
  const [data, setData] = useState(null);
  const [isPending, setIsPending] = useState(false);
  const [error, setError] = useState(null);
  const [checkupResult, setCheckupResult] = useState(null);
  const [payload, setPayload] = useState(null);
  const fRead = new FileReader();
  const [nama_input, setNamaInput] = useState("");
  const [penyakit_input, setPenyakitInput] = useState("");
  const [dna_input, setDnaInput] = useState("");
  const [dna_error, setDnaError] = useState(null);
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
      console.log(payload);
      if (payload) {
        setData(null);
        setIsPending(true);
        setError(null);
        setCheckupResult(null);
        const { data: d, isPending: p, error: err } = await fetchPOST(endpoint, payload);
        setIsPending(p);
        if (err) {
          setError(err);
        } else {
          setData(d);
          setCheckupResult(`${d.data.tanggal_pred} - ${d.data.nama_pasien} - ${d.data.nama_penyakit} - ${d.data.status} - ${d.data.similarity}`);
        }
      }
    };
    doAsync();
  }, [payload]);
  const fileHandler = (e) => {
    fRead.onload = async (ev) => {
      const tdat = ev.target.result;
      setDnaInput(tdat.trim());
    };
    fRead.readAsText(e.target.files[0]);
    setDnaFileName(e.target.files[0].name);
  };
  const handleSubmit = (e) => {
    e.preventDefault();
    setResDisplay(false);
    const trimmedname = nama_input.trim();
    const trimmedpenyakit = penyakit_input.trim();
    const trimmeddna = dna_input.trim();
    setResDisplay(true);
    setDnaError(null);
    setPayload({
      nama_pasien: nama_input.trim(),
      dna_pasien: dna_input.trim(),
      nama_penyakit: penyakit_input.trim(),
    });
  };
  return (
    <Card>
      <CardContent>
        <Grid sx={{ flexGrow: 1, p: 3 }} container columnSpacing={2}>
          <Box component={Grid} item align="center" boxShadow={0} xs={12} sx={{ mb: 5, fontSize: 16 }} fontWeight="fontWeightBold">
            <Typography component="div">
              <Box sx={{ fontSize: "h4.fontSize", fontWeight: "bold" }}>DNA Checkup</Box>
            </Typography>
          </Box>
          <Box component={Grid} item boxShadow={0} xs={12}>
            <Box component="form" onSubmit={handleSubmit}>
              <Grid sx={{ flexGrow: 1 }} container columnSpacing={2}>
                <Box component={Grid} item boxShadow={0} xs={12} md={4} sx={{ mb: 5 }}>
                  <TextField required fullWidth id="inp_nama" onChange={(e) => setNamaInput(e.target.value)} label="Nama Pengguna" />
                </Box>
                <Box component={Grid} item boxShadow={0} xs={12} md={4} sx={{ mb: 5 }}>
                  <Box sx={{ fontSize: 12 }}>DNA Sequence (.txt): {dna_filename}</Box>
                  <Box component={Stack} direction="column" justifyContent="center">
                    <Button variant="contained" component="label">
                      Upload File
                      <input required onChange={fileHandler} type="file" accept=".txt" hidden />
                    </Button>
                  </Box>
                  <Typography component="div" color="secondary">
                    <Box sx={{ fontSize: 12, mt: 1, fontWeight: "bold" }}>{dna_error}</Box>
                  </Typography>
                </Box>
                <Box component={Grid} item boxShadow={0} xs={12} md={4} sx={{ mb: 5 }}>
                  <TextField required fullWidth id="inp_penyakit" label="Penyakit" onChange={(e) => setPenyakitInput(e.target.value)} />
                </Box>
                <Box component={Grid} item boxShadow={0} xs={8} sx={{ mb: 5, display: { xs: "none", md: "block" } }}></Box>
                <Box component={Grid} align="right" item boxShadow={0} xs={12} md={4} sx={{ mb: 5 }}>
                  <Button type="submit" color="secondary" variant="contained">
                    CHECK
                  </Button>
                </Box>
              </Grid>
            </Box>
          </Box>
          <Box component={Grid} item xs={12} sx={{ display: resDisplay ? "block" : "none" }}>
            <Grid container columnSpacing={2}>
              <Box component={Grid} item align="center" boxShadow={0} xs={12} sx={{ mb: 2, fontSize: 16 }} fontWeight="fontWeightBold">
                <Typography component="div">
                  <Box sx={{ fontSize: "h5.fontSize", fontWeight: "bold" }}>Result</Box>
                </Typography>
              </Box>
              <Box component={Grid} item align="center" boxShadow={0} xs={12} sx={{ mb: 5, fontSize: 16 }} fontWeight="fontWeightBold">
                <Typography component="div">
                  <Box sx={{ fontSize: "h5.fontSize", fontWeight: "bold", display: isPending ? "block" : "none" }}>Loading...</Box>
                  <Box sx={[{ display: error ? "block" : "none" }, sxerr]}>Error: {error}</Box>
                  <Box sx={[{ display: data ? "block" : "none" }, sxsuc]}>{checkupResult}</Box>
                </Typography>
              </Box>
              {/* <Box component={Grid} item boxShadow={0} xs={12} xl={3} sx={{mb: 5}} >
                            <TextField
                                fullWidth
                                id="res_tanggal_out"
                                label="Tanggal"
                                defaultValue={res_tanggal}
                                InputProps={{
                                    readOnly: true,
                                }}
                            />
                        </Box>
                        <Box component={Grid} item boxShadow={0} xs={12} xl={2} sx={{mb: 5}} >
                            <TextField
                                fullWidth
                                id="res_pengguna_out"
                                label="Pengguna"
                                defaultValue={res_nama}
                                InputProps={{
                                    readOnly: true,
                                }}
                            />
                        </Box>
                        <Box component={Grid} item boxShadow={0} xs={12} xl={2} sx={{mb: 5}} >
                            <TextField
                                fullWidth
                                id="res_penyakit_out"
                                label="Penyakit"
                                defaultValue={res_penyakit}
                                InputProps={{
                                    readOnly: true,
                                }}
                            />
                        </Box>
                        <Box component={Grid} item boxShadow={0} xs={12} xl={2} sx={{mb: 5}} >
                            <TextField
                                fullWidth
                                id="res_kemiripan_out"
                                label="Kemiripan"
                                defaultValue={res_kemiripan}
                                InputProps={{
                                    readOnly: true,
                                }}
                            />
                        </Box>
                        <Box component={Grid} item boxShadow={0} xs={12} xl={3} sx={{mb: 5}} >
                            <TextField
                                fullWidth
                                id="res_diagnosa_out"
                                label="Diagnosa"
                                defaultValue={res_diagnosa}
                                InputProps={{
                                    readOnly: true,
                                }}
                            />
                        </Box> */}
            </Grid>
          </Box>
        </Grid>
      </CardContent>
    </Card>
  );
}

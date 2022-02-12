import React, { useState } from 'react';
import {
  Button, Container, makeStyles, Paper, Typography,
} from '@material-ui/core';
import { foodService } from '../services/food-service';
import useAsync from '../hooks/use-async';
import FoodContainer from '../components/food-container';
import FoodResponse from '../models/food-response';

const useStyles = makeStyles((theme) => ({
  root: {
    display: 'flex',
    textAlign: 'center',
    flexDirection: 'column',
    width: '1120px',
    margin: '0 auto',
    backgroundColor: theme.palette.background.default,
    minHeight: `calc(100vh - 2*${theme.mixins.toolbar.minHeight}px)`,
    alignItems: 'center',
    padding: 0,
  },
  label: {
    color: theme.palette.primary.dark,
    fontFamily: theme.typography.fontFamily,
    fontSize: '24px',
    fontWeight: 500,
    marginTop: '40px',
    marginBottom: '20px',
  },
  uploadButton: {
    marginBottom: '20px',
  },
}));

export default function SearchPage() {
  const [photo, setPhoto] = useState<File>();

  const classes = useStyles();

  const handleImageChange = (event: React.FormEvent) => {
    event.preventDefault();
    const { files } = event.target as HTMLInputElement;

    if (files && files.length > 0) {
      setPhoto(files[0]);
    }
  };

  const {
    data: foodResponse, loading: loadingFood, error: foodError,
  } = useAsync<FoodResponse>(() => foodService.searchForImage(photo), [photo]);

  return (
    <Paper elevation={3} className={classes.root}>
      <Container>
        <Typography className={classes.label}>
          Upload an image to get information about foods
        </Typography>

        {photo && (
          <div>
            <img alt="not found" width="250px" src={URL.createObjectURL(photo)} />
          </div>
        )}

        <label htmlFor="photo">
          <input
            accept="image/*"
            style={{ display: 'none' }}
            id="photo"
            name="photo"
            type="file"
            multiple={false}
            onChange={handleImageChange}
          />

          <Button
            component="span"
            variant="contained"
            color="primary"
            className={classes.uploadButton}
          >
            Choose Image
          </Button>

        </label>

        <FoodContainer foodResponse={foodResponse} error={foodError} loading={loadingFood} />
      </Container>
    </Paper>
  );
}

import {
  Button, Link, makeStyles, Paper, Typography,
} from '@material-ui/core';
import { Link as RouterLink } from 'react-router-dom';
import React from 'react';

const useStyles = makeStyles((theme) => ({
  paper: {
    display: 'flex',
    flexDirection: 'column',
    width: '58%',
    // 100% height - header and footer
    minHeight: `calc(100vh - 2*${theme.mixins.toolbar.minHeight}px)`,
    textAlign: 'center',
    background: theme.palette.background.paper,
    margin: '0 auto',
    alignItems: 'center',
  },
  title: {
    marginTop: theme.spacing(10),
    marginBottom: theme.spacing(11),
    color: theme.palette.primary.dark,
  },
  button: {
    height: '56px',
    width: '100px',
  },
}));

export default function MainPage() {
  const classes = useStyles();

  return (
    <Paper className={classes.paper}>
      <Typography variant="h2" component="h2" gutterBottom className={classes.title}>
        Find a food in
        <br />
        FoodData Central
      </Typography>
      <Link component={RouterLink} to="/search">
        <Button
          color="primary"
          variant="contained"
          className={classes.button}
        >
          Search
        </Button>
      </Link>
    </Paper>
  );
}

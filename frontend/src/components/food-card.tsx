import React from 'react';
import { Box, makeStyles, Typography } from '@material-ui/core';
import FoodInfo from '../models/food-info';

const useStyles = makeStyles((theme) => ({
  root: {
    boxShadow: '0px 3px 6px #20025D26',
    borderRadius: '4px',
    marginTop: '10px',
    marginLeft: 'auto',
    marginRight: 'auto',
  },
  title: {
    color: theme.palette.primary.dark,
    letterSpacing: '0.25px',
    overflowWrap: 'break-word',
    wordWrap: 'break-word',
    fontSize: '20px',
    width: '100%',
    marginTop: '10px',
    marginBottom: '8px',
  },
  foodDetails: {
    width: '98%',
    marginRight: 'auto',
    paddingLeft: '24px',
    marginBottom: '8px',
    letterSpacing: '0.4px',
    fontSize: '16px',
    overflowWrap: 'break-word',
    wordWrap: 'break-word',
  },
}));

interface FoodCardProps {
  food: FoodInfo;
}

export default function FoodCard({ food }: FoodCardProps) {
  const classes = useStyles();

  return (
    <Box className={classes.root}>
      <Typography className={classes.title} variant="body1">
        <b>{food.description}</b>
      </Typography>
      <Typography className={classes.foodDetails} variant="body1">
        <b>FdcId:</b>
        {' '}
        {food.fdcId}
      </Typography>
      {food.gtinUpc && (
        <Typography className={classes.foodDetails} variant="body1">
          <b>GtinUpc:</b>
          {' '}
          {food.gtinUpc}
        </Typography>
      )}
      {food.ingredients && (
        <Typography className={classes.foodDetails} variant="body1">
          <b>Ingredients:</b>
          {' '}
          {food.ingredients}
        </Typography>
      )}
    </Box>
  );
}

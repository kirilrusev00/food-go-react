import React from 'react';
import { Box, makeStyles } from '@material-ui/core';
import FoodCard from './food-card';
import FoodResponse from '../models/food-response';
import { Loading } from './loading';

const useStyles = makeStyles({
  foodGrid: {
    maxWidth: '95vw',
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
    marginTop: '10px',
    marginBottom: '10px',
  },
});

export interface FoodContainerProps {
  foodResponse: FoodResponse | undefined;
  loading: boolean;
  error: boolean;
}

export default function FoodContainer({ foodResponse, loading, error }: FoodContainerProps) {
  const classes = useStyles();

  return (
    <Loading loading={loading} error={error}>
      {foodResponse && (
        <Box className={classes.foodGrid}>
          {foodResponse?.foods.map((foodInfo) => (
            <FoodCard
              key={foodInfo.fdcId}
              food={foodInfo}
            />
          ))}
        </Box>
      )}
    </Loading>
  );
}

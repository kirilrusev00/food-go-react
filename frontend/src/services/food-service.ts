import FoodResponse from '../models/food-response';
import { httpService } from './http-service';

class FoodService {
  async searchForImage(photo: File | undefined): Promise<FoodResponse> {
    if (!photo) {
      return {
        foods: [],
      };
    }

    const formData = new FormData();
    formData.append('photo', photo);

    const response = await httpService.post<FoodResponse>('/image', formData);
    return response;
  }
}

/* eslint-disable-next-line import/prefer-default-export */
export const foodService = new FoodService();

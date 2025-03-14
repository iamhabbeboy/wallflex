import { imagePathStore } from "../store/app";

export const dispatcher = (value: string) => {
  imagePathStore.set(value);
}


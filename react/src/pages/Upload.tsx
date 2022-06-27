import React, { SyntheticEvent, useEffect, useRef, useState } from "react";
import styles from "../styles/leftMenu.module.sass";
import uploadStyles from "../styles/uploadLayout.module.sass";

export default function Upload(props: { userid: number }) {
  const inputRef = useRef<HTMLInputElement>(null);
  const formRef = useRef<HTMLFormElement>(null);

  const upload = async (event: any) => {
    event.preventDefault();
    let file = null;
    if (inputRef.current?.files) {
      console.log(inputRef?.current?.files[0]);
      file = inputRef?.current?.files[0];
    }
    if (file) {
      const formData = new FormData();
      console.log(file);
      formData.append("cloud", file, file.name);
      console.log(formData);
      const response = await fetch(
        `https://localhost:8000/api/projects/upload/${props.userid}/${file?.name}`,
        {
          method: "POST",
          credentials: "include",
          body: formData,
        }
      );
      return response;
    }
  };

  return (
    <form
      id={"form"}
      ref={formRef}
      className={styles.wrapper}
      onSubmit={(event) => upload(event)}
    >
      <div>
        <label htmlFor="image_uploads">Выберите файл для загрузки</label>
        <input
          className={uploadStyles["custom-file-input"]}
          ref={inputRef}
          type="file"
          id="proj_uploads"
          name="image_uploads"
        />
      </div>
      <div>
        <button type="submit" className={uploadStyles.buttonUpload}>
          Submit
        </button>
      </div>
    </form>
  );
}

colprint() (
   text="${1#'[#'??????']'}"
   r="${1%"$text"}"
   r="${r%']'}"
   r="${r#'[#'}"
   r="${r:-FFFFFF}"
   b="${r#????}"
   r="${r%??}"
   g="${r#??}"
   r="${r%??}"
   printf '\e[38;2;%d;%d;%dm%s\e[0m\n' "0x$r" "0x$g" "0x$b" "$text"
)
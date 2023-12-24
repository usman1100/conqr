# conqr: Conquer Speed Limits with Concurrent Downloads

**Conqr** is a blazing-fast, concurrent CLI download manager written in Go, designed to shatter speed limits and turbocharge your downloads. It achieves this by employing a powerful strategy:

- **Splitting Downloads into Chunks:** Conqr intelligently divides downloads into multiple segments, enabling simultaneous downloads of each chunk.
- **Concurrent Download Power:** By downloading chunks concurrently, Conqr can often bypass speed limits imposed by servers, significantly accelerating the overall download process.
- **Seamless Stitching:** Once all chunks have been downloaded, Conqr expertly reassembles them into the complete file, ensuring a smooth and efficient experience.

**Key Features:**

- **Concurrent Downloads:** Turbocharge downloads by downloading multiple chunks simultaneously.
- **Speed Limit Circumvention:** Often bypass speed limits imposed by servers, maximizing download speeds.
- **Easy-to-Use CLI:** Simple and intuitive command-line interface for effortless downloads.
- **Written in Go:** Leverages the efficiency and performance of the Go language.

**Usage:**

1. To download a file:

   ```bash
   conqr <URL>
   ```

   Replace `<URL>` with the actual URL of the file you wish to download.

2. To specify the number of chunks:

   ```bash
   conqr -c <number_of_chunks> <URL>
   ```

**Contributing:**

We welcome contributions! Feel free to open issues or pull requests.

**License:**

This project is licensed under the MIT License.

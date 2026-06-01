function change(s, prog, version) {
  const phoneMatch = s.match(/Phone:\s*([^\n]+)/);
  const versionMatch = s.match(/Version:\s*([^\n]+)/);

  const phone = phoneMatch ? phoneMatch[1].trim() : "";
  const ver = versionMatch ? versionMatch[1].trim() : "";

  const validPhone = /^\+1-\d{3}-\d{3}-\d{4}$/.test(phone);
  const validVersion = /^\d+\.\d+$/.test(ver);

  if (!validPhone || !validVersion) {
    return "ERROR: VERSION or PHONE";
  }

  const newVersion = ver === "2.0" ? "2.0" : version;

  const lines = s.split("\n");
  const out = [];

  for (const line of lines) {
    if (/Corporation/.test(line) || /Level/.test(line)) {
      continue;
    } else if (/Program\s+[Tt]itle/.test(line)) {
      out.push("Program: " + prog);
    } else if (/^Author:/.test(line)) {
      out.push("Author: g964");
    } else if (/^Phone:/.test(line)) {
      out.push("Phone: +1-503-555-0090");
    } else if (/^Date:/.test(line)) {
      out.push("Date: 2019-01-01");
    } else if (/^Version:/.test(line)) {
      out.push("Version: " + newVersion);
    } else {
      out.push(line);
    }
  }

  return out.join(" ");
}
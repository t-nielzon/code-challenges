/*
You have two arguments: string - a string of random letters(only lowercase)
and array - an array of strings(feelings). Your task is to return how many
specific feelings are in the array.

Each letter in string participates in the formation of all feelings.
If the feeling can be formed once - plus one to the answer.
If the feeling can be formed several times from different letters - plus one to the answer.

Examples:
  string -> 'yliausoenvjw', array -> ['anger', 'awe', 'joy', 'love', 'grief']
  output -> '3 feelings.' // 'awe', 'joy', 'love'

  string -> 'griefgriefgrief', array -> ['anger', 'awe', 'joy', 'love', 'grief']
  output -> '1 feeling.' // 'grief'

  string -> 'abcdkasdfvkadf', array -> ['desire', 'joy', 'shame', 'longing', 'fear']
  output -> '0 feelings.'
*/

function countFeelings(string, array) {
}
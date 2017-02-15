/*Package rand provides random-generation methods for types that are not available in the standard rand package
 */
package rand

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

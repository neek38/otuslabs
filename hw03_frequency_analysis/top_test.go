package hw03_frequency_analysis //nolint:golint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Change to true if needed
var taskWithAsteriskIsCompleted = true

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

var animalsText = `Живо́тные (лат. Animalia) — традиционно (со времён Аристотеля) выделяемая категория организмов, 
	 	в настоящее время рассматривается в качестве биологического царства. 
Животные являются основным объектом изучения зоологии.
Животные относятся к эукариотам (в клетках имеются ядра). 	Классическими признаками животных считаются: 
гетеротрофность (питание готовыми 
	     органическими соединениями) и   способность активно 	передвигаться. Впрочем, существует немало животных
 , ведущих неподвижный образ жизни, а гетеротрофность свойственна грибам и некоторым растениям-паразитам.
Русское слово «животное»   образовано от «живот», в прошлом означавшего «жизнь, 	имущество». 
 В быту под терминами «дикие животные», «домашние животные» часто понимаются только млекопитающие
или четвероногие наземные позвоночные (млекопитающие, пресмыкающиеся и земноводные). Однако в науке за
термином «животные» закреплено более широкое значение, соответствующее латинскому Animalia (см. выше).
В научном смысле к животным,         помимо млекопитающих, пресмыкающихся и земноводных,
относится огромное   множество других организмов: рыбы, птицы, 	
насекомые, паукообразные, моллюски, морские звёзды, черви и другие.
При этом ранее к этому царству относили
многих гетеротрофных протистов и делили животных на подцарства: одноклеточные Protozoa и многоклеточные Metazoa. 
  Сейчас название «животные» в таксономическом 	смысле закрепилось за многоклеточными.
В таком понимании животные как таксон 		
имеют более определённые признаки — для них   	характерны оогамия, многотканево́е строение, наличие как
минимум двух зародышевых листков, стадий бластулы и гаструлы в зародышевом развитии. Человек относится к царству 
животных, но традиционно изучается отдельно. У   подавляющего большинства животных есть мышцы и нервы, а не имеющие их 
группы — губки, пластинчатые, мезозои, книдоспоридии — возможно, лишились их вторично.
В то же время, в науке термин «животные» иногда предлагается использовать и в ещё более широком значении, 
   подразумевая под животными не таксон, а тип организации — жизненную форму, основанную на подвижности.
В настоящее время (Zhang, 2013) учёными описано более 1,6 млн видов животных (включая более 133 тыс. ископаемых видов;
Zhang, 2013), большинство из которых составляют членистоногие (более 1,3 млн видов, 78 %), 
моллюски (более 118 тыс. видов) и позвоночные (более 42 тыс. видов)`

var germanText = `Gestern ging ich mit meinem Sohn und meiner Frau 
in den Tierpark. Wir fuhren 	zuerst mit der Eisenbahn bis an den Hauptbahnhof. 
Anschließend bestiegen wir die Straßenbahn und fuhren bis an die 
Haltestelle    Zoologischer Garten.
Der Tierpark befindet 		sich auf einem kleinen Berg oben und wir
konnten die wunderbare Aussicht auf die Stadt und den See genießen.
Vor der Kasse mussten wir 	15 Minuten anstehen, weil an diesem Sonntag viele 
Parkbesucher anwesend waren. Nachdem wir den Eintritt bezahlt hatten, empfingen uns bereits 



die ersten Tiere, Königspinguine wohlgemerkt. Es war 
Fütterungszeit und die Wärter warfen tote Fische in das Pinguingehege.
Der Tierpark ist sehr weitläufig und wir trafen im Elefantenpark ein.
Vor ein paar Wochen ist ein Elefantenbaby auf die Welt gekommen. Der Kleine 
lief friedlich umher, fraß immer und wälzte sich im Sand. Nun waren die Affen 
an der Reihe und wir konnten die verschiedenen Arten 
(Schimpansen,  Gorillas, Orang-Utans, 	usw.) bestaunen.
Die Schimpansen schwangen sich ganz locker von Liane zu Liane, während 
die Gorillas Müde	 auf dem Boden lagen. Am Nachmittag sahen wir weitere Tiere, 
vom Känguru, Zebra,	 Löwen bis zu den Delfinen, stand alles auf dem Programm. 
Glücklich und zufrieden  	 verließen wir den Tierpark und fuhren wieder nach Hause.`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		assert.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{"он", "а", "и", "что", "ты", "не", "если", "то", "его", "кристофер", "робин", "в"}
			assert.Subset(t, expected, Top10(text))
		} else {
			expected := []string{"он", "и", "а", "что", "ты", "не", "если", "-", "то", "Кристофер"}
			assert.ElementsMatch(t, expected, Top10(text))
		}
	})

	t.Run("positive test with in russian", func(t *testing.T) {
		expected := []string{"в", "и", "более", "животные", "животных", "видов", "к", "время", "тыс", "а"}
		assert.Subset(t, expected, Top10(animalsText))
	})

	t.Run("positive test in german", func(t *testing.T) {
		expected := []string{"die", "und", "wir", "der", "den", "auf", "tierpark", "an", "sich", "fuhren", "bis", "ein"}
		assert.Subset(t, expected, Top10(germanText))
	})
}

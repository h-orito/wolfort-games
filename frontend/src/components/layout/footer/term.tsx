export default function Term() {
  return (
    <div className='text-xs leading-5'>
      <p className='my-4'>
        この利用規約（以下、「本規約」といいます。）は、wolfort
        games管理人（以下、「管理人」といいます。）がこのウェブサイト上で提供するサービス（以下、「本サービス」といいます。）の利用条件を定めるものです。登録ユーザーの皆さま（以下、「ユーザー」といいます。）には、本規約に従って、本サービスをご利用いただきます。
      </p>
      <Label>第1条（適用）</Label>
      <OrderdList>
        <ListItem>
          本規約は、ユーザーと管理人との間の本サービスの利用に関わる一切の関係に適用されるものとします。
        </ListItem>
        <ListItem>
          管理人は本サービスに関し、本規約のほか、ご利用にあたってのルール等、各種の定め（以下、「個別規定」といいます。）をすることがあります。これら個別規定はその名称のいかんに関わらず、本規約の一部を構成するものとします。
        </ListItem>
        <ListItem>
          本規約の規定が前条の個別規定の規定と矛盾する場合には、個別規定において特段の定めなき限り、個別規定の規定が優先されるものとします。
        </ListItem>
      </OrderdList>
      <Label>第2条（利用登録）</Label>
      <OrderdList>
        <ListItem>
          本サービスにおいては、登録希望者が本規約に同意の上、管理人の定める方法によって利用登録を申請し、管理人がこれを承認することによって、利用登録が完了するものとします。
        </ListItem>
        <ListItem>
          管理人は、利用登録の申請者に以下の事由があると判断した場合、利用登録の申請を承認しないことがあり、その理由については一切の開示義務を負わないものとします。
          <OrderdList className='pl-4'>
            <ListItem>利用登録の申請に際して虚偽の事項を届け出た場合</ListItem>
            <ListItem>
              本規約に違反したことがある者からの申請である場合
            </ListItem>
            <ListItem>
              その他、管理人が利用登録を相当でないと判断した場合
            </ListItem>
          </OrderdList>
        </ListItem>
      </OrderdList>
      <Label>第3条（ユーザーIDおよびパスワードの管理）</Label>
      <OrderdList>
        <ListItem>
          ユーザーは、自己の責任において、本サービスのユーザーIDおよびパスワードを適切に管理するものとします。
        </ListItem>
        <ListItem>
          ユーザーは、いかなる場合にも、ユーザーIDおよびパスワードを第三者に譲渡または貸与し、もしくは第三者と共用することはできません。管理人は、ユーザーIDとパスワードの組み合わせが登録情報と一致してログインされた場合には、そのユーザーIDを登録しているユーザー自身による利用とみなします。
        </ListItem>
        <ListItem>
          ユーザーID及びパスワードが第三者によって使用されたことによって生じた損害は、管理人に故意又は重大な過失がある場合を除き、管理人は一切の責任を負わないものとします。
        </ListItem>
      </OrderdList>
      <Label>第4条（利用料金および支払方法）</Label>
      <OrderdList>
        <ListItem>
          ユーザーは、本サービスの有料部分の対価として、管理人が別途定め、本ウェブサイトに表示する利用料金を、管理人が指定する方法により支払うものとします（2023年8月時点では、有料部分は存在しません。）。
        </ListItem>
        <ListItem>
          ユーザーが利用料金の支払を遅滞した場合には、ユーザーは年14.6％の割合による遅延損害金を支払うものとします。
        </ListItem>
      </OrderdList>
      <Label>第5条（禁止事項）</Label>
      <p className='mt-4'>
        ユーザーは、本サービスの利用にあたり、以下の行為をしてはなりません。
      </p>
      <OrderdList>
        <ListItem>法令または公序良俗に違反する行為</ListItem>
        <ListItem>犯罪行為に関連する行為</ListItem>
        <ListItem>
          本サービスの内容等、本サービスに含まれる著作権、商標権ほか知的財産権を侵害する行為
        </ListItem>
        <ListItem>
          管理人、ほかのユーザー、またはその他第三者のサーバーまたはネットワークの機能を破壊したり、妨害したりする行為
        </ListItem>
        <ListItem>
          本サービスによって得られた情報を商業的に利用する行為
        </ListItem>
        <ListItem>管理人のサービスの運営を妨害するおそれのある行為</ListItem>
        <ListItem>不正アクセスをし、またはこれを試みる行為</ListItem>
        <ListItem>
          他のユーザーに関する個人情報等を収集または蓄積する行為
        </ListItem>
        <ListItem>不正な目的を持って本サービスを利用する行為</ListItem>
        <ListItem>
          本サービスの他のユーザーまたはその他の第三者に不利益、損害、不快感を与える行為
        </ListItem>
        <ListItem>他のユーザーに成りすます行為</ListItem>
        <ListItem>
          管理人が許諾しない本サービス上での宣伝、広告、勧誘、または営業行為
        </ListItem>
        <ListItem>面識のない異性との出会いを目的とした行為</ListItem>
        <ListItem>
          管理人のサービスに関連して、反社会的勢力に対して直接または間接に利益を供与する行為
        </ListItem>
        <ListItem>その他、管理人が不適切と判断する行為</ListItem>
      </OrderdList>
      <Label>第6条（本サービスの提供の停止等）</Label>
      <OrderdList>
        <ListItem>
          管理人は、以下のいずれかの事由があると判断した場合、ユーザーに事前に通知することなく本サービスの全部または一部の提供を停止または中断することができるものとします。
          <OrderdList className='pl-4'>
            <ListItem>
              本サービスにかかるコンピュータシステムの保守点検または更新を行う場合
            </ListItem>
            <ListItem>
              地震、落雷、火災、停電または天災などの不可抗力により、本サービスの提供が困難となった場合
            </ListItem>
            <ListItem>
              コンピュータまたは通信回線等が事故により停止した場合
            </ListItem>
            <ListItem>
              その他、管理人が本サービスの提供が困難と判断した場合
            </ListItem>
          </OrderdList>
        </ListItem>
        <ListItem>
          管理人は、本サービスの提供の停止または中断により、ユーザーまたは第三者が被ったいかなる不利益または損害についても、一切の責任を負わないものとします。
        </ListItem>
      </OrderdList>
      <Label>第7条（利用制限および登録抹消）</Label>
      <OrderdList>
        <ListItem>
          管理人は、ユーザーが以下のいずれかに該当する場合には、事前の通知なく、ユーザーに対して、本サービスの全部もしくは一部の利用を制限し、またはユーザーとしての登録を抹消することができるものとします。
          <ListItem>本規約のいずれかの条項に違反した場合</ListItem>
          <ListItem>登録事項に虚偽の事実があることが判明した場合</ListItem>
          <ListItem>料金等の支払債務の不履行があった場合</ListItem>
          <ListItem>管理人からの連絡に対し、一定期間返答がない場合</ListItem>
          <ListItem>
            本サービスについて、最終の利用から一定期間利用がない場合
          </ListItem>
          <ListItem>
            その他、管理人が本サービスの利用を適当でないと判断した場合
          </ListItem>
        </ListItem>
        <ListItem>
          管理人は、本条に基づき管理人が行った行為によりユーザーに生じた損害について、一切の責任を負いません。
        </ListItem>
      </OrderdList>
      <Label>第8条（退会）</Label>
      <p className='my-4'>
        ユーザーは、管理人の定める退会手続により、本サービスから退会できるものとします。
      </p>
      <Label>第9条（保証の否認および免責事項）</Label>
      <OrderdList>
        <ListItem>
          管理人は、本サービスに事実上または法律上の瑕疵（安全性、信頼性、正確性、完全性、有効性、特定の目的への適合性、セキュリティなどに関する欠陥、エラーやバグ、権利侵害などを含みます。）がないことを明示的にも黙示的にも保証しておりません。
        </ListItem>
        <ListItem>
          管理人は、本サービスに起因してユーザーに生じたあらゆる損害について、管理人の故意又は重過失による場合を除き、一切の責任を負いません。ただし、本サービスに関する管理人とユーザーとの間の契約（本規約を含みます。）が消費者契約法に定める消費者契約となる場合、この免責規定は適用されません。
        </ListItem>
        <ListItem>
          前項ただし書に定める場合であっても、管理人は、管理人の過失（重過失を除きます。）による債務不履行または不法行為によりユーザーに生じた損害のうち特別な事情から生じた損害（管理人またはユーザーが損害発生につき予見し、または予見し得た場合を含みます。）について一切の責任を負いません。また、管理人の過失（重過失を除きます。）による債務不履行または不法行為によりユーザーに生じた損害の賠償は、ユーザーから当該損害が発生した月に受領した利用料の額を上限とします。
        </ListItem>
        <ListItem>
          管理人は、本サービスに関して、ユーザーと他のユーザーまたは第三者との間において生じた取引、連絡または紛争等について一切責任を負いません。
        </ListItem>
      </OrderdList>
      <Label>第10条（サービス内容の変更等）</Label>
      <p className='my-4'>
        管理人は、ユーザーへの事前の告知なしに、本サービスの内容を変更、追加または廃止することがあり、ユーザーはこれを承諾するものとします。
      </p>
      <Label>第11条（利用規約の変更）</Label>
      <p className='my-4'>
        管理人は、ユーザーの個別の同意を要せず、本規約を変更することができるものとします。
      </p>
      <Label>第12条（個人情報の取扱い）</Label>
      <p className='my-4'>
        管理人は、本サービスの利用によって取得する個人情報については、管理人「プライバシーポリシー」に従い適切に取り扱うものとします。
      </p>
      <Label>第13条（通知または連絡）</Label>
      <p className='my-4'>
        ユーザーと管理人との間の通知または連絡は、管理人の定める方法によって行うものとします。管理人は、ユーザーから、管理人が別途定める方式に従った変更届け出がない限り、現在登録されている連絡先が有効なものとみなして当該連絡先へ通知または連絡を行い、これらは、発信時にユーザーへ到達したものとみなします。
      </p>
      <Label>第14条（権利義務の譲渡の禁止）</Label>
      <p className='my-4'>
        ユーザーは、管理人の書面による事前の承諾なく、利用契約上の地位または本規約に基づく権利もしくは義務を第三者に譲渡し、または担保に供することはできません。
      </p>
      <Label>第15条（準拠法・裁判管轄）</Label>
      <OrderdList>
        <ListItem>本規約の解釈にあたっては、日本法を準拠法とします。</ListItem>
        <ListItem>
          本サービスに関して紛争が生じた場合には、管理人の在住所在地を管轄する裁判所を専属的合意管轄とします。
        </ListItem>
      </OrderdList>
      <p className='mt-6 flex justify-end'>以上</p>
    </div>
  )
}

const Label = (props: { children: React.ReactNode }) => (
  <label className='my-4 text-sm font-bold'>{props.children}</label>
)

const OrderdList = (props: {
  children: React.ReactNode
  className?: string
}) => (
  <ol
    className={`${
      props.className ?? ''
    } list-inside list-decimal py-2 text-left text-xs`}
  >
    {props.children}
  </ol>
)

const ListItem = (props: { children: React.ReactNode; className?: string }) => (
  <li className={`${props.className ?? ''} my-2`}>{props.children}</li>
)
